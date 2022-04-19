import '../App.css'
import React, { useEffect } from 'react'
import { get, post } from '../common/api'
import { updateArray, objectMap } from '../common/misc'
import { PlayerList, INum, UserList } from '../common/input'
import { Stack, InputNumber, Divider, Checkbox, CheckboxGroup, Panel, PanelGroup, Radio, RadioGroup } from 'rsuite'
import Header from '../navbar/navbar'
import { playerValue, playersArray } from '../common/players'
import { useParams, useNavigate } from 'react-router-dom'
import { setMinutes } from 'rsuite/esm/utils/dateUtils'
import ArrowRightIcon from '@rsuite/icons/ArrowRight'

const MainZi = ({ idx, setFu }) => {
    const [number, setNumber] = React.useState('0')
    const [isMing, setIsMing] = React.useState(false)
    const [isYaoJiu, setIsYaoJiu] = React.useState(false)
    const [len, setLen] = React.useState(0)

    const title = '面子' + (idx + 1).toString()

    function updateFu() {
        if (number == '3') {
            switch (len) {
                case 0:
                    setFu(2)
                    break
                case 1:
                    setFu(4)
                    break
                case 2:
                    setFu(8)
                    break
            }
        } else if (number == '4') {
            switch (len) {
                case 0:
                    setFu(8)
                    break
                case 1:
                    setFu(16)
                    break
                case 2:
                    setFu(32)
                    break
            }
        } else {
            setFu(0)
        }
    }

    useEffect(() => {
        updateFu()
    }, [number, len]
    )

    return (
        <Panel header={title} bordered>
            <RadioGroup
                inline
                name='mainzi3or4'
                value={number}
                onChange={value => {
                    setNumber(value)
                }}
            >
                <Radio value='0'>顺子</Radio>
                <Radio value='3'>刻子</Radio>
                <Radio value='4'>杠子</Radio>
            </RadioGroup>
            {number > 2 && (
                <CheckboxGroup inline onChange={v => setLen(v.length)}>
                    <Checkbox value='yaojiu'> 一九</Checkbox>
                    {((number == 3)
                        ? <Checkbox value='mingan'> 暗刻</Checkbox>
                        : <Checkbox value='mingan'> 暗杠</Checkbox>)}
                </CheckboxGroup>
            )}
        </Panel>
    )
}

const QueTou = ({ setValue }) => {
    const [ziFeng, setZiFeng] = React.useState(false)
    const [changFeng, setChangFeng] = React.useState(false)
    const [sanYuan, setSanYuan] = React.useState(false)

    useEffect(() => {
        if (sanYuan) {
            setValue(2)
        } else {
            if (ziFeng && changFeng) {
                setValue(4)
            } else if (ziFeng || changFeng) {
                setValue(2)
            } else {
                setValue(0)
            }
        }
    }, [ziFeng, changFeng, sanYuan]
    )

    return (
        <Panel header='雀头'>
            <Checkbox
                value='ziFeng' inline checked={ziFeng} onChange={(_, v) => {
                    setSanYuan(false)
                    setZiFeng(v)
                }}
            > 自风
            </Checkbox>
            <Checkbox
                value='changFeng' inline checked={changFeng} onChange={(_, v) => {
                    setSanYuan(false)
                    setChangFeng(v)
                }}
            > 场风
            </Checkbox>
            <Checkbox
                value='sanYuan' inline checked={sanYuan} onChange={(_, v) => {
                    setZiFeng(false)
                    setChangFeng(false)
                    setSanYuan(v)
                }}
            > 三元
            </Checkbox>

        </Panel>
    )
}

function roundUpNearest10(num) {
    return Math.ceil(num / 10) * 10
}

function roundUpNearest100(num) {
    return Math.ceil(num / 100) * 100
}

const Fu = ({ setValue, isZiMo }) => {
    const [menQian, setMenQian] = React.useState(false)
    const [yiMainTing, setYiMainTing] = React.useState(false)
    const [queTou, setQueTou] = React.useState(0)
    const [mainziFu, setMainziFu] = React.useState(Array(4).fill(0))

    const updateMainziFu = (fu, index) => {
        // console.log('fu:' + fu + ',idx:' + index)
        setMainziFu((prev) => prev.map((el, i) => (i !== index ? el : fu)))
    }

    useEffect(() => {
        setValue(20 +
            roundUpNearest10(
                (menQian && !isZiMo ? 10 : 0) +
                (isZiMo ? 2 : 0) +
                (yiMainTing ? 2 : 0) +
                Number(queTou) +
                mainziFu.reduce((partialSum, a) => partialSum + Number(a), 0))
        )
    })

    return (
        <>
            <div className='mt-5 flex flex-col items-center min-h-screen py-2'>
                <Panel header='杂'>
                    <Checkbox
                        value='menQian' inline checked={menQian} onChange={(_, v) => { setMenQian(v) }}
                    > 门前清
                    </Checkbox>
                    <Checkbox
                        value='yiMainTing' inline checked={yiMainTing} onChange={(_, v) => { setYiMainTing(v) }}
                    > 一面听
                    </Checkbox>
                </Panel>
                <QueTou setValue={setQueTou} />
                <br />
                <PanelGroup>
                    {mainziFu.map((el, i) => (
                        <MainZi key={'mainzi' + i} idx={i} setFu={x => updateMainziFu(x, i)} />
                    ))}
                </PanelGroup>
            </div>
        </>
    )
}

const JMPoints = () => {
    const [fu, setFu] = React.useState(20)
    const [fan, setFan] = React.useState(1)
    const [minFan, setMinFan] = React.useState(1)
    const [isZhuang, setIsZhuang] = React.useState(false)
    const [isZiMo, setIsZiMo] = React.useState(false)
    const [isQiDui, setIsQiDui] = React.useState(false)
    const [isZiMoPingHu, setIsZiMoPingHu] = React.useState(false)
    const [lianZhuang, setLianZhuang] = React.useState(0)
    const [res, setRes] = React.useState('0')

    function updateRes() {
        let base = 0
        switch (Number(fan)) {
            case 5:
                base = 2000
                break
            case 6:
            case 7:
                base = 3000
                break
            case 8:
            case 9:
            case 10:
                base = 4000
                break
            case 11:
            case 12:
                base = 6000
                break
            case 13:
                base = 8000
                break
            default:
                base = fu * Math.pow(2, 2 + Number(fan))
                if (base > 2000) {
                    base = 2000
                }
        }
        switch (true) {
            case isZhuang && isZiMo:
                setRes(roundUpNearest100(2 * base) + lianZhuang * 300)
                break
            case isZhuang && !isZiMo:
                setRes(roundUpNearest100(6 * base) + lianZhuang * 300)
                break
            case !isZhuang && isZiMo:
                setRes((roundUpNearest100(2 * base) + lianZhuang * 100) + ', ' + (roundUpNearest100(base) + lianZhuang * 100))
                break
            case !isZhuang && !isZiMo:
                setRes(roundUpNearest100(4 * base) + lianZhuang * 300)
                break
        }
    }

    useEffect(() => {
        updateRes()
    })

    useEffect(() => {
        if (isZiMoPingHu) {
            setFu(20)
            setIsQiDui(false)
            setIsZiMo(true)
            setMinFan(2)
            if (fan < 2) {
                setFan(2)
            }
        } else {
            setMinFan(1)
            if (fan == 2) {
                setFan(1)
            }
        }
    }, [isZiMoPingHu])
    return (
        <>
            <Header />
            <div class='mt-5 min-h-screen py-2'>
                <div class='flex flex-col items-center'>
                    <Panel bordered shaded>
                        {isZhuang ? '庄家' : ' 闲家'}{isZiMo ? '自摸' : '荣和'} :  {fan} 番 {fan < 5 ? fu + '符' : ''}  <ArrowRightIcon style={{ fontSize: '2em', marginBottom: '3px' }} /> {res}
                    </Panel>
                    <Panel header='结果'>
                        <Checkbox
                            value='isZhuang' inline checked={isZhuang} onChange={(_, v) => { setIsZhuang(v) }}
                        > 庄家和
                        </Checkbox>
                        <Checkbox
                            value='isZiMo' inline checked={isZiMo} onChange={(_, v) => {
                                if (!v && isZiMoPingHu) {
                                    setIsZiMoPingHu(false)
                                }
                                setIsZiMo(v)
                            }}
                        > 自摸
                        </Checkbox>
                        <br />
                        <br />
                        <Stack spacing={6}>
                            <InputNumber prefix='番' type='number' value={fan} onChange={(n, _) => setFan(n)} min={minFan} />
                            <InputNumber prefix='连庄' type='number' value={lianZhuang} onChange={(n, _) => setLianZhuang(n)} min={0} />
                        </Stack>
                        <br />
                        <Panel header='特殊情况' collapsible bordered>
                            <Checkbox
                                value='isQiDui' inline checked={isQiDui} onChange={(_, v) => {
                                    if (v) {
                                        setFu(25)
                                        setIsZiMoPingHu(false)
                                    }
                                    setIsQiDui(v)
                                }}
                            > 七对
                            </Checkbox>
                            <Checkbox
                                value='isZiMoPingHu' inline checked={isZiMoPingHu} onChange={(_, v) => {
                                    setIsZiMoPingHu(v)
                                }}
                            > 自摸平和
                            </Checkbox>
                        </Panel>
                    </Panel>
                </div>
                {!isQiDui && !isZiMoPingHu && fan < 5 &&
                    <>
                        <Divider> 符</Divider>
                        <Fu setValue={setFu} isZiMo={isZiMo} />
                    </>}

            </div>
        </>
    )
}

export { JMPoints }
