# VkVideoEncodeH264PictureInfoKHR(3) Manual Page

## Name

VkVideoEncodeH264PictureInfoKHR - Structure specifies H.264 encode frame
parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264PictureInfoKHR {
    VkStructureType                             sType;
    const void*                                 pNext;
    uint32_t                                    naluSliceEntryCount;
    const VkVideoEncodeH264NaluSliceInfoKHR*    pNaluSliceEntries;
    const StdVideoEncodeH264PictureInfo*        pStdPictureInfo;
    VkBool32                                    generatePrefixNalu;
} VkVideoEncodeH264PictureInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `naluSliceEntryCount` is the number of elements in
  `pNaluSliceEntries`.

- `pNaluSliceEntries` is a pointer to an array of `naluSliceEntryCount`
  [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)
  structures specifying the parameters of the individual H.264 slices to
  encode for the input picture.

- `pStdPictureInfo` is a pointer to a `StdVideoEncodeH264PictureInfo`
  structure specifying <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-picture-info"
  target="_blank" rel="noopener">H.264 picture information</a>.

- `generatePrefixNalu` controls whether prefix NALUs are generated
  before slice NALUs into the target bitstream, as defined in sections
  7.3.2.12 and 7.4.2.12 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>.

## <a href="#_description" class="anchor"></a>Description

This structure is specified in the `pNext` chain of the
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html) structure passed to
[vkCmdEncodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEncodeVideoKHR.html) to specify the
codec-specific picture information for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264"
target="_blank" rel="noopener">H.264 encode operation</a>.

Encode Input Picture Information  
When this structure is specified in the `pNext` chain of the
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html) structure passed to
[vkCmdEncodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEncodeVideoKHR.html), the information related
to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-input-picture-info"
target="_blank" rel="noopener">encode input picture</a> is defined as
follows:

- The image subregion used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-picture-data-access"
  target="_blank" rel="noopener">H.264 Encode Picture Data Access</a>
  section.

- The encode input picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-picture-info"
  target="_blank" rel="noopener">H.264 picture information</a> provided
  in `pStdPictureInfo`.

<!-- -->

Std Picture Information  
The members of the `StdVideoEncodeH264PictureInfo` structure pointed to
by `pStdPictureInfo` are interpreted as follows:

- `flags.reserved` and `reserved1` are used only for padding purposes
  and are otherwise ignored;

- `flags.IdrPicFlag` as defined in section 7.4.1 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- `flags.is_reference` as defined in section 3.136 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- `seq_parameter_set_id` and `pic_parameter_set_id` are used to identify
  the active parameter sets, as described below;

- `primary_pic_type` as defined in section 7.4.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- `PicOrderCnt` as defined in section 8.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- `temporal_id` as defined in section G.7.4.1.1 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- if `pRefLists` is not `NULL`, then it is a pointer to a
  `StdVideoEncodeH264ReferenceListsInfo` structure that is interpreted
  as follows:

  - `flags.reserved` is used only for padding purposes and is otherwise
    ignored;

  - `ref_pic_list_modification_flag_l0` and
    `ref_pic_list_modification_flag_l1` as defined in section 7.4.3.1 of
    the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
    target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

  - `num_ref_idx_l0_active_minus1` and `num_ref_idx_l1_active_minus1` as
    defined in section 7.4.3 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
    target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

  - `RefPicList0` and `RefPicList1` as defined in section 8.2.4 of the
    <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
    target="_blank" rel="noopener">ITU-T H.264 Specification</a> where
    each element of these arrays either identifies an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-active-reference-picture-info"
    target="_blank" rel="noopener">active reference picture</a> using
    its <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
    target="_blank" rel="noopener">DPB slot</a> index or contains the
    value `STD_VIDEO_H264_NO_REFERENCE_PICTURE` to indicate “no
    reference picture”;

  - if `refList0ModOpCount` is not zero, then `pRefList0ModOperations`
    is a pointer to an array of `refList0ModOpCount` number of
    `StdVideoEncodeH264RefListModEntry` structures specifying the
    modification parameters for the reference list L0 as defined in
    section 7.4.3.1 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
    target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

  - if `refList1ModOpCount` is not zero, then `pRefList1ModOperations`
    is a pointer to an array of `refList1ModOpCount` number of
    `StdVideoEncodeH264RefListModEntry` structures specifying the
    modification parameters for the reference list L1 as defined in
    section 7.4.3.1 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
    target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

  - if `refPicMarkingOpCount` is not zero, then
    `refPicMarkingOperations` is a pointer to an array of
    `refPicMarkingOpCount` number of
    `StdVideoEncodeH264RefPicMarkingEntry` structures specifying the
    reference picture marking parameters as defined in section 7.4.3.3
    of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
    target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- all other members are interpreted as defined in section 7.4.3 of the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>.

Reference picture setup is controlled by the value of
`StdVideoEncodeH264PictureInfo`::`flags.is_reference`. If it is set and
a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is specified,
then the latter is used as the target of picture reconstruction to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">activate</a> the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
target="_blank" rel="noopener">DPB slot</a> specified in
`pEncodeInfo->pSetupReferenceSlot->slotIndex`. If
`StdVideoEncodeH264PictureInfo`::`flags.is_reference` is not set, but a
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is specified,
then the corresponding picture reference associated with the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
target="_blank" rel="noopener">DPB slot</a> is invalidated, as described
in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">DPB Slot States</a> section.

Active Parameter Sets  
The members of the `StdVideoEncodeH264PictureInfo` structure pointed to
by `pStdPictureInfo` are used to select the active parameter sets to use
from the bound video session parameters object, as follows:

- <span id="encode-h264-active-sps"></span> The *active SPS* is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-sps"
  target="_blank" rel="noopener">SPS</a> identified by the key specified
  in `StdVideoEncodeH264PictureInfo`::`seq_parameter_set_id`.

- <span id="encode-h264-active-pps"></span> The *active PPS* is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-pps"
  target="_blank" rel="noopener">PPS</a> identified by the key specified
  by the pair constructed from
  `StdVideoEncodeH264PictureInfo`::`seq_parameter_set_id` and
  `StdVideoEncodeH264PictureInfo`::`pic_parameter_set_id`.

H.264 encoding uses *explicit weighted sample prediction* for a slice,
as defined in section 8.4.2.3 of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
target="_blank" rel="noopener">ITU-T H.264 Specification</a>, if any of
the following conditions are true for the active <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-pps"
target="_blank" rel="noopener">PPS</a> and the `pStdSliceHeader` member
of the corresponding element of `pNaluSliceEntries`:

- `pStdSliceHeader->slice_type` is `STD_VIDEO_H264_SLICE_TYPE_P` and
  `weighted_pred_flag` is enabled in the active PPS.

- `pStdSliceHeader->slice_type` is `STD_VIDEO_H264_SLICE_TYPE_B` and
  `weighted_bipred_idc` in the active PPS equals
  `STD_VIDEO_H264_WEIGHTED_BIPRED_IDC_EXPLICIT`.

Valid Usage

- <a
  href="#VUID-VkVideoEncodeH264PictureInfoKHR-naluSliceEntryCount-08301"
  id="VUID-VkVideoEncodeH264PictureInfoKHR-naluSliceEntryCount-08301"></a>
  VUID-VkVideoEncodeH264PictureInfoKHR-naluSliceEntryCount-08301  
  `naluSliceEntryCount` **must** be between `1` and
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`maxSliceCount`,
  inclusive, as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile

- <a href="#VUID-VkVideoEncodeH264PictureInfoKHR-flags-08304"
  id="VUID-VkVideoEncodeH264PictureInfoKHR-flags-08304"></a>
  VUID-VkVideoEncodeH264PictureInfoKHR-flags-08304  
  If
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`flags`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile, does not include
  `VK_VIDEO_ENCODE_H264_CAPABILITY_GENERATE_PREFIX_NALU_BIT_KHR`, then
  `generatePrefixNalu` **must** be `VK_FALSE`

- <a href="#VUID-VkVideoEncodeH264PictureInfoKHR-flags-08314"
  id="VUID-VkVideoEncodeH264PictureInfoKHR-flags-08314"></a>
  VUID-VkVideoEncodeH264PictureInfoKHR-flags-08314  
  If
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`flags`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile, does not include
  `VK_VIDEO_ENCODE_H264_CAPABILITY_PREDICTION_WEIGHT_TABLE_GENERATED_BIT_KHR`
  and the slice corresponding to any element of `pNaluSliceEntries` uses
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-weighted-pred"
  target="_blank" rel="noopener">explicit weighted sample prediction</a>,
  then
  [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)::`pStdSliceHeader->pWeightTable`
  **must** not be `NULL` for that element of `pNaluSliceEntries`

- <a href="#VUID-VkVideoEncodeH264PictureInfoKHR-flags-08315"
  id="VUID-VkVideoEncodeH264PictureInfoKHR-flags-08315"></a>
  VUID-VkVideoEncodeH264PictureInfoKHR-flags-08315  
  If
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`flags`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile, does not include
  `VK_VIDEO_ENCODE_H264_CAPABILITY_DIFFERENT_SLICE_TYPE_BIT_KHR`, then
  [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)::`pStdSliceHeader->slice_type`
  **must** be identical for all elements of `pNaluSliceEntries`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH264PictureInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH264PictureInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH264PictureInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_PICTURE_INFO_KHR`

- <a
  href="#VUID-VkVideoEncodeH264PictureInfoKHR-pNaluSliceEntries-parameter"
  id="VUID-VkVideoEncodeH264PictureInfoKHR-pNaluSliceEntries-parameter"></a>
  VUID-VkVideoEncodeH264PictureInfoKHR-pNaluSliceEntries-parameter  
  `pNaluSliceEntries` **must** be a valid pointer to an array of
  `naluSliceEntryCount` valid
  [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)
  structures

- <a
  href="#VUID-VkVideoEncodeH264PictureInfoKHR-pStdPictureInfo-parameter"
  id="VUID-VkVideoEncodeH264PictureInfoKHR-pStdPictureInfo-parameter"></a>
  VUID-VkVideoEncodeH264PictureInfoKHR-pStdPictureInfo-parameter  
  `pStdPictureInfo` **must** be a valid pointer to a valid
  `StdVideoEncodeH264PictureInfo` value

- <a
  href="#VUID-VkVideoEncodeH264PictureInfoKHR-naluSliceEntryCount-arraylength"
  id="VUID-VkVideoEncodeH264PictureInfoKHR-naluSliceEntryCount-arraylength"></a>
  VUID-VkVideoEncodeH264PictureInfoKHR-naluSliceEntryCount-arraylength  
  `naluSliceEntryCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h264.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH264PictureInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
