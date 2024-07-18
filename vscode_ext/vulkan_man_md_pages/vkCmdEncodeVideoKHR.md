# vkCmdEncodeVideoKHR(3) Manual Page

## Name

vkCmdEncodeVideoKHR - Launch video encode operations



## <a href="#_c_specification" class="anchor"></a>C Specification

To launch video encode operations, call:

``` c
// Provided by VK_KHR_video_encode_queue
void vkCmdEncodeVideoKHR(
    VkCommandBuffer                             commandBuffer,
    const VkVideoEncodeInfoKHR*                 pEncodeInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to record the command.

- `pEncodeInfo` is a pointer to a
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html) structure specifying
  the parameters of the video encode operations.

## <a href="#_description" class="anchor"></a>Description

Each call issues one or more video encode operations. The implicit
parameter `opCount` corresponds to the number of video encode operations
issued by the command. After calling this command, the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active-query-index"
target="_blank" rel="noopener">active query index</a> of each <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
target="_blank" rel="noopener">active</a> query is incremented by
`opCount`.

Currently each call to this command results in the issue of a single
video encode operation.

If the bound video session was created with
`VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR` and the `pNext` chain
of `pEncodeInfo` includes a
[VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure
with its `queryPool` member specifying a valid `VkQueryPool` handle,
then this command will execute a query for each video encode operation
issued by it.

Active Reference Picture Information  
The list of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#active-reference-pictures"
target="_blank" rel="noopener">active reference pictures</a> used by a
video encode operation is a list of image subregions used as the source
of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-picture"
target="_blank" rel="noopener">reference picture</a> data and related
parameters, and is derived from the
[VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
structures provided as the elements of the
`pEncodeInfo->pReferenceSlots` array. For each element of
`pEncodeInfo->pReferenceSlots`, one or more elements are added to the
active reference picture list, as defined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-codec-specific-semantics"
target="_blank" rel="noopener">codec-specific semantics</a>. Each
element of this list contains the following information:

- The image subregion within the image subresource <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">referred</a> to by the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
  target="_blank" rel="noopener">video picture resource</a> used as the
  reference picture.

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> index the reference
  picture is associated with.

- The codec-specific reference information related to the reference
  picture.

<!-- -->

Reconstructed Picture Information  
Information related to the optional <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reconstructed-picture"
target="_blank" rel="noopener">reconstructed picture</a> used by a video
encode operation is derived from the
[VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
structure pointed to by `pEncodeInfo->pSetupReferenceSlot`, if not
`NULL`, as defined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-codec-specific-semantics"
target="_blank" rel="noopener">codec-specific semantics</a>, and
consists of the following:

- The image subregion within the image subresource <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">referred</a> to by the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
  target="_blank" rel="noopener">video picture resource</a> used as the
  reconstructed picture.

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> index to use for picture
  reconstruction.

- The codec-specific reference information related to the reconstructed
  picture.

Specifying a valid
[VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
structure in `pEncodeInfo->pSetupReferenceSlot` is always required,
unless the video session was created with
[VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlot`
equal to zero. However, the DPB slot identified by
`pEncodeInfo->pSetupReferenceSlot->slotIndex` is only <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">activated</a> with the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reconstructed-picture"
target="_blank" rel="noopener">reconstructed picture</a> specified in
`pEncodeInfo->pSetupReferenceSlot->pPictureResource` if reference
picture setup is requested according to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-codec-specific-semantics"
target="_blank" rel="noopener">codec-specific semantics</a>.

If reconstructed picture information is specified, but reference picture
setup is not requested, according to the codec-specific semantics, the
contents of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
target="_blank" rel="noopener">video picture resource</a> corresponding
to the reconstructed picture will be undefined after the video encode
operation.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Some implementations may always output the reconstructed picture or
use it as temporary storage during the video encode operation even when
the reconstructed picture is not marked for future reference.</p></td>
</tr>
</tbody>
</table>

Encode Input Picture Information  
Information related to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-input-picture"
target="_blank" rel="noopener">encode input picture</a> used by a video
encode operation is derived from `pEncodeInfo->srcPictureResource` and
any codec-specific parameters provided in the `pEncodeInfo->pNext`
chain, as defined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-codec-specific-semantics"
target="_blank" rel="noopener">codec-specific semantics</a>, and
consists of the following:

- The image subregion within the image subresource <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">referred</a> to by the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
  target="_blank" rel="noopener">video picture resource</a> used as the
  encode input picture.

- The codec-specific picture information related to the encoded picture.

Several limiting values are defined below that are referenced by the
relevant valid usage statements of this command.

- Let `uint32_t activeReferencePictureCount` be the size of the list of
  active reference pictures used by the video encode operation. Unless
  otherwise defined, `activeReferencePictureCount` is set to the value
  of `pEncodeInfo->referenceSlotCount`.

- Let `VkOffset2D codedOffsetGranularity` be the minimum alignment
  requirement for the coded offset of video picture resources. Unless
  otherwise defined, the value of the `x` and `y` members of
  `codedOffsetGranularity` are `0`.

- Let `uint32_t dpbFrameUseCount[]` be an array of size `maxDpbSlots`,
  where `maxDpbSlots` is the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlots`
  the bound video session was created with, with each element indicating
  the number of times a frame associated with the corresponding DPB slot
  index is referred to by the video coding operation. Let the initial
  value of each element of the array be `0`.

  - If `pEncodeInfo->pSetupReferenceSlot` is not `NULL`, then
    `dpbFrameUseCount[i]` is incremented by one, where `i` equals
    `pEncodeInfo->pSetupReferenceSlot->slotIndex`.

  - For each element of `pEncodeInfo->pReferenceSlots`,
    `dpbFrameUseCount[i]` is incremented by one, where `i` equals the
    `slotIndex` member of the corresponding element.

- Let `VkExtent2D maxCodingBlockSize` be the maximum codec-specific
  coding block size that **may** be used by the video encode operation.

  - If the bound video session object was created with an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-profile"
    target="_blank" rel="noopener">H.264 encode profile</a>, then let
    `maxCodingBlockSize` be equal to the size of an H.264 macroblock,
    i.e. `{16,16}`.

  - If the bound video session object was created with an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-profile"
    target="_blank" rel="noopener">H.265 encode profile</a>, then let
    `maxCodingBlockSize` be equal to the maximum H.265 coding block size
    that **may** be used by the video encode operation derived as the
    maximum of the CTB sizes corresponding to the
    [VkVideoEncodeH265CtbSizeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CtbSizeFlagBitsKHR.html)
    bits set in
    [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`ctbSizes`,
    as returned by
    [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
    for the video profile the bound video session was created with.

  - Otherwise, `maxCodingBlockSize` is undefined.

- If `maxCodingBlockSize` is defined, then let
  `VkExtent2D minCodingBlockExtent` be the coded extent of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-input-picture"
  target="_blank" rel="noopener">encode input picture</a> expressed in
  terms of codec-specific coding blocks, assuming the maximum size of
  such coding blocks, as defined by `maxCodingBlockSize`, calculated
  from the value of the `codedExtent` member of
  `pEncodeInfo->srcPictureResource` as follows:

  - `minCodingBlockExtent.width` = (`codedExtent.width`  
    `maxCodingBlockSize.width` - 1) / `maxCodingBlockSize.width`

  - `minCodingBlockExtent.height` = (`codedExtent.height`  
    `maxCodingBlockSize.height` - 1) / `maxCodingBlockSize.height`

- If the bound video session object was created with an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-profile"
  target="_blank" rel="noopener">H.264 encode profile</a>, then:

  - Let `StdVideoH264PictureType h264PictureType` be the picture type of
    the encoded picture set to the value of
    `pStdPictureInfo->primary_pic_type` specified in the
    [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
    structure included in the `pEncodeInfo->pNext` chain.

  - Let `StdVideoH264PictureType h264L0PictureTypes[]` and
    `StdVideoH264PictureType h264L1PictureTypes[]` be the picture types
    of the reference pictures in the L0 and L1 reference lists,
    respectively. If `pStdPictureInfo->pRefLists` specified in the
    [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
    structure included in the `pEncodeInfo->pNext` chain is not `NULL`,
    then for each reference index specified in the elements of the
    `pStdPictureInfo->pRefLists->RefPicList0` and
    `pStdPictureInfo->pRefLists->RefPicList1` arrays, if the reference
    index is not `STD_VIDEO_H264_NO_REFERENCE_PICTURE`,
    `pStdReferenceInfo->primary_pic_type` is added to
    `h264L0PictureTypes` or `h264L1PictureTypes`, respectively, where
    `pStdReferenceInfo` is the member of the
    [VkVideoEncodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264DpbSlotInfoKHR.html)
    structure included in the `pNext` chain of the element of
    `pEncodeInfo->pReferenceSlots` for which `slotIndex` equals the
    reference index in question.

- If the bound video session object was created with an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-profile"
  target="_blank" rel="noopener">H.265 encode profile</a>, then:

  - Let `StdVideoH265PictureType h265PictureType` be the picture type of
    the encoded picture set to the value of `pStdPictureInfo->pic_type`
    specified in the
    [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
    structure included in the `pEncodeInfo->pNext` chain.

  - Let `StdVideoH265PictureType h265L0PictureTypes[]` and
    `StdVideoH265PictureType h265L1PictureTypes[]` be the picture types
    of the reference pictures in the L0 and L1 reference lists,
    respectively. If `pStdPictureInfo->pRefLists` specified in the
    [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
    structure included in the `pEncodeInfo->pNext` chain is not `NULL`,
    then for each reference index specified in the elements of the
    `pStdPictureInfo->pRefLists->RefPicList0` and
    `pStdPictureInfo->pRefLists->RefPicList1` arrays, if the reference
    index is not `STD_VIDEO_H265_NO_REFERENCE_PICTURE`,
    `pStdReferenceInfo->pic_type` is added to `h265L0PictureTypes` or
    `h265L1PictureTypes`, respectively, where `pStdReferenceInfo` is the
    member of the
    [VkVideoEncodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265DpbSlotInfoKHR.html)
    structure included in the `pNext` chain of the element of
    `pEncodeInfo->pReferenceSlots` for which `slotIndex` equals the
    reference index in question.

Valid Usage

- <a href="#VUID-vkCmdEncodeVideoKHR-None-08250"
  id="VUID-vkCmdEncodeVideoKHR-None-08250"></a>
  VUID-vkCmdEncodeVideoKHR-None-08250  
  The bound video session **must** have been created with an encode
  operation

- <a href="#VUID-vkCmdEncodeVideoKHR-None-07012"
  id="VUID-vkCmdEncodeVideoKHR-None-07012"></a>
  VUID-vkCmdEncodeVideoKHR-None-07012  
  The bound video session **must** not be in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-uninitialized"
  target="_blank" rel="noopener">uninitialized</a> state at the time the
  command is executed on the device

- <a href="#VUID-vkCmdEncodeVideoKHR-None-08318"
  id="VUID-vkCmdEncodeVideoKHR-None-08318"></a>
  VUID-vkCmdEncodeVideoKHR-None-08318  
  The bound video session parameters object **must** have been created
  with the currently set <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-quality-level"
  target="_blank" rel="noopener">video encode quality level</a> for the
  bound video session at the time the command is executed on the device

- <a href="#VUID-vkCmdEncodeVideoKHR-opCount-07174"
  id="VUID-vkCmdEncodeVideoKHR-opCount-07174"></a>
  VUID-vkCmdEncodeVideoKHR-opCount-07174  
  For each <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a> query, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active-query-index"
  target="_blank" rel="noopener">active query index</a> corresponding to
  the query type of that query plus `opCount` **must** be less than or
  equal to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-last-activatable-query-index"
  target="_blank" rel="noopener">last activatable query index</a>
  corresponding to the query type of that query plus one

- <a href="#VUID-vkCmdEncodeVideoKHR-pNext-08360"
  id="VUID-vkCmdEncodeVideoKHR-pNext-08360"></a>
  VUID-vkCmdEncodeVideoKHR-pNext-08360  
  If the bound video session was created with
  `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`, and the `pNext`
  chain of `pEncodeInfo` includes a
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure
  with its `queryPool` member specifying a valid `VkQueryPool` handle,
  then
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html)::queryCount
  **must** equal `opCount`

- <a href="#VUID-vkCmdEncodeVideoKHR-pNext-08361"
  id="VUID-vkCmdEncodeVideoKHR-pNext-08361"></a>
  VUID-vkCmdEncodeVideoKHR-pNext-08361  
  If the bound video session was created with
  `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`, and the `pNext`
  chain of `pEncodeInfo` includes a
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure
  with its `queryPool` member specifying a valid `VkQueryPool` handle,
  then all the queries used by the command, as specified by the
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure,
  **must** be *unavailable*

- <a href="#VUID-vkCmdEncodeVideoKHR-queryType-08362"
  id="VUID-vkCmdEncodeVideoKHR-queryType-08362"></a>
  VUID-vkCmdEncodeVideoKHR-queryType-08362  
  If the bound video session was created with
  `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`, then the `queryType`
  used to create the `queryPool` specified in the
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure
  included in the `pNext` chain of `pEncodeInfo` **must** be
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` or
  `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`

- <a href="#VUID-vkCmdEncodeVideoKHR-queryPool-08363"
  id="VUID-vkCmdEncodeVideoKHR-queryPool-08363"></a>
  VUID-vkCmdEncodeVideoKHR-queryPool-08363  
  If the bound video session was created with
  `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`, then the `queryPool`
  specified in the
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure
  included in the `pNext` chain of `pEncodeInfo` **must** have been
  created with a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)
  structure included in the `pNext` chain of
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html) identical to the
  one specified in
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`pVideoProfile`
  the bound video session was created with

- <a href="#VUID-vkCmdEncodeVideoKHR-queryType-08364"
  id="VUID-vkCmdEncodeVideoKHR-queryType-08364"></a>
  VUID-vkCmdEncodeVideoKHR-queryType-08364  
  If the bound video session was created with
  `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`, and the `queryType`
  used to create the `queryPool` specified in the
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure
  included in the `pNext` chain of `pEncodeInfo` is
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`, then the `VkCommandPool` that
  `commandBuffer` was allocated from **must** have been created with a
  queue family index that supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-result-status-only"
  target="_blank" rel="noopener">result status queries</a>, as indicated
  by
  [VkQueueFamilyQueryResultStatusPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyQueryResultStatusPropertiesKHR.html)::`queryResultStatusSupport`

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08201"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08201"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08201  
  `pEncodeInfo->dstBuffer` **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profile-compatibility"
  target="_blank" rel="noopener">compatible</a> with the video profile
  the bound video session was created with

- <a href="#VUID-vkCmdEncodeVideoKHR-commandBuffer-08202"
  id="VUID-vkCmdEncodeVideoKHR-commandBuffer-08202"></a>
  VUID-vkCmdEncodeVideoKHR-commandBuffer-08202  
  If `commandBuffer` is an unprotected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, then `pEncodeInfo->dstBuffer` **must** not be a
  protected buffer

- <a href="#VUID-vkCmdEncodeVideoKHR-commandBuffer-08203"
  id="VUID-vkCmdEncodeVideoKHR-commandBuffer-08203"></a>
  VUID-vkCmdEncodeVideoKHR-commandBuffer-08203  
  If `commandBuffer` is a protected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, then `pEncodeInfo->dstBuffer` **must** be a protected
  buffer

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08204"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08204"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08204  
  `pEncodeInfo->dstBufferOffset` **must** be an integer multiple of
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`minBitstreamBufferOffsetAlignment`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08205"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08205"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08205  
  `pEncodeInfo->dstBufferRange` **must** be an integer multiple of
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`minBitstreamBufferSizeAlignment`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08206"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08206"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08206  
  `pEncodeInfo->srcPictureResource.imageViewBinding` **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profile-compatibility"
  target="_blank" rel="noopener">compatible</a> with the video profile
  the bound video session was created with

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08207"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08207"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08207  
  The format of `pEncodeInfo->srcPictureResource.imageViewBinding`
  **must** match the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`pictureFormat`
  the bound video session was created with

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08208"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08208"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08208  
  `pEncodeInfo->srcPictureResource.codedOffset` **must** be an integer
  multiple of `codedOffsetGranularity`

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08209"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08209"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08209  
  `pEncodeInfo->srcPictureResource.codedExtent` **must** be between
  `minCodedExtent` and `maxCodedExtent`, inclusive, the bound video
  session was created with

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08210"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08210"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08210  
  `pEncodeInfo->srcPictureResource.imageViewBinding` **must** have been
  created with `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`

- <a href="#VUID-vkCmdEncodeVideoKHR-commandBuffer-08211"
  id="VUID-vkCmdEncodeVideoKHR-commandBuffer-08211"></a>
  VUID-vkCmdEncodeVideoKHR-commandBuffer-08211  
  If `commandBuffer` is an unprotected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, then `pEncodeInfo->srcPictureResource.imageViewBinding`
  **must** not have been created from a protected image

- <a href="#VUID-vkCmdEncodeVideoKHR-commandBuffer-08212"
  id="VUID-vkCmdEncodeVideoKHR-commandBuffer-08212"></a>
  VUID-vkCmdEncodeVideoKHR-commandBuffer-08212  
  If `commandBuffer` is a protected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, then `pEncodeInfo->srcPictureResource.imageViewBinding`
  **must** have been created from a protected image

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08377"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08377"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08377  
  `pEncodeInfo->pSetupReferenceSlot` **must** not be `NULL` unless the
  bound video session was created with
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlots`
  equal to zero

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08213"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08213"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08213  
  If `pEncodeInfo->pSetupReferenceSlot` is not `NULL`, then
  `pEncodeInfo->pSetupReferenceSlot->slotIndex` **must** be less than
  the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlots`
  specified when the bound video session was created

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08214"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08214"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08214  
  If `pEncodeInfo->pSetupReferenceSlot` is not `NULL`, then
  `pEncodeInfo->pSetupReferenceSlot->pPictureResource->codedOffset`
  **must** be an integer multiple of `codedOffsetGranularity`

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08215"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08215"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08215  
  If `pEncodeInfo->pSetupReferenceSlot` is not `NULL`, then
  `pEncodeInfo->pSetupReferenceSlot->pPictureResource` **must** <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-matching"
  target="_blank" rel="noopener">match</a> one of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#bound-reference-picture-resources"
  target="_blank" rel="noopener">bound reference picture resource</a>

- <a href="#VUID-vkCmdEncodeVideoKHR-activeReferencePictureCount-08216"
  id="VUID-vkCmdEncodeVideoKHR-activeReferencePictureCount-08216"></a>
  VUID-vkCmdEncodeVideoKHR-activeReferencePictureCount-08216  
  `activeReferencePictureCount` **must** be less than or equal to the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxActiveReferencePictures`
  specified when the bound video session was created

- <a href="#VUID-vkCmdEncodeVideoKHR-slotIndex-08217"
  id="VUID-vkCmdEncodeVideoKHR-slotIndex-08217"></a>
  VUID-vkCmdEncodeVideoKHR-slotIndex-08217  
  The `slotIndex` member of each element of
  `pEncodeInfo->pReferenceSlots` **must** be less than the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlots`
  specified when the bound video session was created

- <a href="#VUID-vkCmdEncodeVideoKHR-codedOffset-08218"
  id="VUID-vkCmdEncodeVideoKHR-codedOffset-08218"></a>
  VUID-vkCmdEncodeVideoKHR-codedOffset-08218  
  The `codedOffset` member of the
  [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html)
  structure pointed to by the `pPictureResource` member of each element
  of `pEncodeInfo->pReferenceSlots` **must** be an integer multiple of
  `codedOffsetGranularity`

- <a href="#VUID-vkCmdEncodeVideoKHR-pPictureResource-08219"
  id="VUID-vkCmdEncodeVideoKHR-pPictureResource-08219"></a>
  VUID-vkCmdEncodeVideoKHR-pPictureResource-08219  
  The `pPictureResource` member of each element of
  `pEncodeInfo->pReferenceSlots` **must** <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-matching"
  target="_blank" rel="noopener">match</a> one of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#bound-reference-picture-resources"
  target="_blank" rel="noopener">bound reference picture resource</a>
  associated with the DPB slot index specified in the `slotIndex` member
  of that element

- <a href="#VUID-vkCmdEncodeVideoKHR-pPictureResource-08220"
  id="VUID-vkCmdEncodeVideoKHR-pPictureResource-08220"></a>
  VUID-vkCmdEncodeVideoKHR-pPictureResource-08220  
  Each video picture resource corresponding to the `pPictureResource`
  member specified in the elements of `pEncodeInfo->pReferenceSlots`
  **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-uniqueness"
  target="_blank" rel="noopener">unique</a> within
  `pEncodeInfo->pReferenceSlots`

- <a href="#VUID-vkCmdEncodeVideoKHR-dpbFrameUseCount-08221"
  id="VUID-vkCmdEncodeVideoKHR-dpbFrameUseCount-08221"></a>
  VUID-vkCmdEncodeVideoKHR-dpbFrameUseCount-08221  
  All elements of `dpbFrameUseCount` **must** be less than or equal to
  `1`

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08222"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08222"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08222  
  The image subresource <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">referred</a> to by
  `pEncodeInfo->srcPictureResource` **must** be in the
  `VK_IMAGE_LAYOUT_VIDEO_ENCODE_SRC_KHR` layout at the time the video
  encode operation is executed on the device

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08223"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08223"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08223  
  If `pEncodeInfo->pSetupReferenceSlot` is not `NULL`, then the image
  subresource <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">referred</a> to by
  `pEncodeInfo->pSetupReferenceSlot->pPictureResource` **must** be in
  the `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DPB_KHR` layout at the time the
  video encode operation is executed on the device

- <a href="#VUID-vkCmdEncodeVideoKHR-pPictureResource-08224"
  id="VUID-vkCmdEncodeVideoKHR-pPictureResource-08224"></a>
  VUID-vkCmdEncodeVideoKHR-pPictureResource-08224  
  The image subresource <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">referred</a> to by the
  `pPictureResource` member of each element of
  `pEncodeInfo->pReferenceSlots` **must** be in the
  `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DPB_KHR` layout at the time the video
  encode operation is executed on the device

- <a href="#VUID-vkCmdEncodeVideoKHR-pNext-08225"
  id="VUID-vkCmdEncodeVideoKHR-pNext-08225"></a>
  VUID-vkCmdEncodeVideoKHR-pNext-08225  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the `pNext` chain
  of `pEncodeInfo` **must** include a
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
  structure

- <a
  href="#VUID-vkCmdEncodeVideoKHR-StdVideoH264SequenceParameterSet-08226"
  id="VUID-vkCmdEncodeVideoKHR-StdVideoH264SequenceParameterSet-08226"></a>
  VUID-vkCmdEncodeVideoKHR-StdVideoH264SequenceParameterSet-08226  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the bound video
  session parameters object **must** contain a
  `StdVideoH264SequenceParameterSet` entry with `seq_parameter_set_id`
  matching `StdVideoEncodeH264PictureInfo`::`seq_parameter_set_id` that
  is provided in the `pStdPictureInfo` member of the
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo`

- <a
  href="#VUID-vkCmdEncodeVideoKHR-StdVideoH264PictureParameterSet-08227"
  id="VUID-vkCmdEncodeVideoKHR-StdVideoH264PictureParameterSet-08227"></a>
  VUID-vkCmdEncodeVideoKHR-StdVideoH264PictureParameterSet-08227  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the bound video
  session parameters object **must** contain a
  `StdVideoH264PictureParameterSet` entry with `seq_parameter_set_id`
  and `pic_parameter_set_id` matching
  `StdVideoEncodeH264PictureInfo`::`seq_parameter_set_id` and
  `StdVideoEncodeH264PictureInfo`::`pic_parameter_set_id`, respectively,
  that are provided in the `pStdPictureInfo` member of the
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo`

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08228"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08228"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08228  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and
  `pEncodeInfo->pSetupReferenceSlot` is not `NULL`, then the `pNext`
  chain of `pEncodeInfo->pSetupReferenceSlot` **must** include a
  [VkVideoEncodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264DpbSlotInfoKHR.html)
  structure

- <a href="#VUID-vkCmdEncodeVideoKHR-pNext-08229"
  id="VUID-vkCmdEncodeVideoKHR-pNext-08229"></a>
  VUID-vkCmdEncodeVideoKHR-pNext-08229  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the `pNext` chain
  of each element of `pEncodeInfo->pReferenceSlots` **must** include a
  [VkVideoEncodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264DpbSlotInfoKHR.html)
  structure

- <a href="#VUID-vkCmdEncodeVideoKHR-constantQp-08269"
  id="VUID-vkCmdEncodeVideoKHR-constantQp-08269"></a>
  VUID-vkCmdEncodeVideoKHR-constantQp-08269  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and the current <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control mode</a> is not
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`, then
  [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)::`constantQp`
  **must** be zero for each element of the `pNaluSliceEntries` member of
  the
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo`

- <a href="#VUID-vkCmdEncodeVideoKHR-constantQp-08270"
  id="VUID-vkCmdEncodeVideoKHR-constantQp-08270"></a>
  VUID-vkCmdEncodeVideoKHR-constantQp-08270  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and the current <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control mode</a> is
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`, then
  [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)::`constantQp`
  **must** be between
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`minQp`
  and
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`maxQp`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, for
  each element of the `pNaluSliceEntries` member of the
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo`

- <a href="#VUID-vkCmdEncodeVideoKHR-constantQp-08271"
  id="VUID-vkCmdEncodeVideoKHR-constantQp-08271"></a>
  VUID-vkCmdEncodeVideoKHR-constantQp-08271  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`flags`
  does not include
  `VK_VIDEO_ENCODE_H264_CAPABILITY_PER_SLICE_CONSTANT_QP_BIT_KHR`, as
  returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)::`constantQp`
  **must** have the same value for each element of the
  `pNaluSliceEntries` member of the
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo`

- <a href="#VUID-vkCmdEncodeVideoKHR-naluSliceEntryCount-08302"
  id="VUID-vkCmdEncodeVideoKHR-naluSliceEntryCount-08302"></a>
  VUID-vkCmdEncodeVideoKHR-naluSliceEntryCount-08302  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the
  `naluSliceEntryCount` member of the
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo` **must** be
  less than or equal to `minCodingBlockExtent.width` multiplied by
  `minCodingBlockExtent.height`

- <a href="#VUID-vkCmdEncodeVideoKHR-naluSliceEntryCount-08312"
  id="VUID-vkCmdEncodeVideoKHR-naluSliceEntryCount-08312"></a>
  VUID-vkCmdEncodeVideoKHR-naluSliceEntryCount-08312  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`flags`
  does not include
  `VK_VIDEO_ENCODE_H264_CAPABILITY_ROW_UNALIGNED_SLICE_BIT_KHR`, as
  returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  the `naluSliceEntryCount` member of the
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo` **must** be
  less than or equal to `minCodingBlockExtent.height`

- <a href="#VUID-vkCmdEncodeVideoKHR-pNext-08352"
  id="VUID-vkCmdEncodeVideoKHR-pNext-08352"></a>
  VUID-vkCmdEncodeVideoKHR-pNext-08352  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, the `pNext` chain of
  `pEncodeInfo` includes a
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
  structure, and `pEncodeInfo->referenceSlotCount` is greater than zero,
  then
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)::`pStdPictureInfo->pRefLists`
  **must** not be `NULL`

- <a href="#VUID-vkCmdEncodeVideoKHR-pNext-08339"
  id="VUID-vkCmdEncodeVideoKHR-pNext-08339"></a>
  VUID-vkCmdEncodeVideoKHR-pNext-08339  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, the `pNext` chain of
  `pEncodeInfo` includes a
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
  structure, and
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)::`pStdPictureInfo->pRefLists`
  is not `NULL`, then each element of the `RefPicList0` and
  `RefPicList1` array members of the
  `StdVideoEncodeH264ReferenceListsInfo` structure pointed to by
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)::`pStdPictureInfo->pRefLists`
  **must** either be `STD_VIDEO_H264_NO_REFERENCE_PICTURE` or **must**
  equal the `slotIndex` member of one of the elements of
  `pEncodeInfo->pReferenceSlots`

- <a href="#VUID-vkCmdEncodeVideoKHR-pNext-08353"
  id="VUID-vkCmdEncodeVideoKHR-pNext-08353"></a>
  VUID-vkCmdEncodeVideoKHR-pNext-08353  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, the `pNext` chain of
  `pEncodeInfo` includes a
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
  structure, and `pEncodeInfo->referenceSlotCount` is greater than zero,
  then the `slotIndex` member of each element of
  `pEncodeInfo->pReferenceSlots` **must** equal one of the elements of
  the `RefPicList0` or `RefPicList1` array members of the
  `StdVideoEncodeH264ReferenceListsInfo` structure pointed to by
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)::`pStdPictureInfo->pRefLists`

- <a href="#VUID-vkCmdEncodeVideoKHR-maxPPictureL0ReferenceCount-08340"
  id="VUID-vkCmdEncodeVideoKHR-maxPPictureL0ReferenceCount-08340"></a>
  VUID-vkCmdEncodeVideoKHR-maxPPictureL0ReferenceCount-08340  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`maxPPictureL0ReferenceCount`
  is zero, as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  `h264PictureType` and each element of `h264L0PictureTypes` and
  `h264L1PictureTypes` **must** not be `STD_VIDEO_H264_PICTURE_TYPE_P`

- <a href="#VUID-vkCmdEncodeVideoKHR-maxBPictureL0ReferenceCount-08341"
  id="VUID-vkCmdEncodeVideoKHR-maxBPictureL0ReferenceCount-08341"></a>
  VUID-vkCmdEncodeVideoKHR-maxBPictureL0ReferenceCount-08341  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`maxBPictureL0ReferenceCount`
  and
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`maxL1ReferenceCount`
  are both zero, as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  `h264PictureType` and each element of `h264L0PictureTypes` and
  `h264L1PictureTypes` **must** not be `STD_VIDEO_H264_PICTURE_TYPE_B`

- <a href="#VUID-vkCmdEncodeVideoKHR-flags-08342"
  id="VUID-vkCmdEncodeVideoKHR-flags-08342"></a>
  VUID-vkCmdEncodeVideoKHR-flags-08342  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`flags`
  does not include
  `VK_VIDEO_ENCODE_H264_CAPABILITY_B_FRAME_IN_L0_LIST_BIT_KHR`, as
  returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  each element of `h264L0PictureTypes` **must** not be
  `STD_VIDEO_H264_PICTURE_TYPE_B`

- <a href="#VUID-vkCmdEncodeVideoKHR-flags-08343"
  id="VUID-vkCmdEncodeVideoKHR-flags-08343"></a>
  VUID-vkCmdEncodeVideoKHR-flags-08343  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`flags`
  does not include
  `VK_VIDEO_ENCODE_H264_CAPABILITY_B_FRAME_IN_L1_LIST_BIT_KHR`, as
  returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  each element of `h264L1PictureTypes` **must** not be
  `STD_VIDEO_H264_PICTURE_TYPE_B`

- <a href="#VUID-vkCmdEncodeVideoKHR-pNext-08230"
  id="VUID-vkCmdEncodeVideoKHR-pNext-08230"></a>
  VUID-vkCmdEncodeVideoKHR-pNext-08230  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the `pNext` chain
  of `pEncodeInfo` **must** include a
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
  structure

- <a href="#VUID-vkCmdEncodeVideoKHR-StdVideoH265VideoParameterSet-08231"
  id="VUID-vkCmdEncodeVideoKHR-StdVideoH265VideoParameterSet-08231"></a>
  VUID-vkCmdEncodeVideoKHR-StdVideoH265VideoParameterSet-08231  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the bound video
  session parameters object **must** contain a
  `StdVideoH265VideoParameterSet` entry with
  `vps_video_parameter_set_id` matching
  `StdVideoEncodeH265PictureInfo`::`sps_video_parameter_set_id` that is
  provided in the `pStdPictureInfo` member of the
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo`

- <a
  href="#VUID-vkCmdEncodeVideoKHR-StdVideoH265SequenceParameterSet-08232"
  id="VUID-vkCmdEncodeVideoKHR-StdVideoH265SequenceParameterSet-08232"></a>
  VUID-vkCmdEncodeVideoKHR-StdVideoH265SequenceParameterSet-08232  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the bound video
  session parameters object **must** contain a
  `StdVideoH265SequenceParameterSet` entry with
  `sps_video_parameter_set_id` and `sps_seq_parameter_set_id` matching
  `StdVideoEncodeH265PictureInfo`::`sps_video_parameter_set_id` and
  `StdVideoEncodeH265PictureInfo`::`pps_seq_parameter_set_id`,
  respectively, that are provided in the `pStdPictureInfo` member of the
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo`

- <a
  href="#VUID-vkCmdEncodeVideoKHR-StdVideoH265PictureParameterSet-08233"
  id="VUID-vkCmdEncodeVideoKHR-StdVideoH265PictureParameterSet-08233"></a>
  VUID-vkCmdEncodeVideoKHR-StdVideoH265PictureParameterSet-08233  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the bound video
  session parameters object **must** contain a
  `StdVideoH265PictureParameterSet` entry with
  `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and
  `pps_pic_parameter_set_id` matching
  `StdVideoEncodeH265PictureInfo`::`sps_video_parameter_set_id`,
  `StdVideoEncodeH265PictureInfo`::`pps_seq_parameter_set_id`, and
  `StdVideoEncodeH265PictureInfo`::`pps_pic_parameter_set_id`,
  respectively, that are provided in the `pStdPictureInfo` member of the
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo`

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08234"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08234"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-08234  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and
  `pEncodeInfo->pSetupReferenceSlot` is not `NULL`, then the `pNext`
  chain of `pEncodeInfo->pSetupReferenceSlot` **must** include a
  [VkVideoEncodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265DpbSlotInfoKHR.html)
  structure

- <a href="#VUID-vkCmdEncodeVideoKHR-pNext-08235"
  id="VUID-vkCmdEncodeVideoKHR-pNext-08235"></a>
  VUID-vkCmdEncodeVideoKHR-pNext-08235  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the `pNext` chain
  of each element of `pEncodeInfo->pReferenceSlots` **must** include a
  [VkVideoEncodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265DpbSlotInfoKHR.html)
  structure

- <a href="#VUID-vkCmdEncodeVideoKHR-constantQp-08272"
  id="VUID-vkCmdEncodeVideoKHR-constantQp-08272"></a>
  VUID-vkCmdEncodeVideoKHR-constantQp-08272  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the current <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control mode</a> is not
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`, then
  [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)::`constantQp`
  **must** be zero for each element of the `pNaluSliceSegmentEntries`
  member of the
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo`

- <a href="#VUID-vkCmdEncodeVideoKHR-constantQp-08273"
  id="VUID-vkCmdEncodeVideoKHR-constantQp-08273"></a>
  VUID-vkCmdEncodeVideoKHR-constantQp-08273  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the current <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control mode</a> is
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`, then
  [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)::`constantQp`
  **must** be between
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`minQp`
  and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxQp`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, for
  each element of the `pNaluSliceSegmentEntries` member of the
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo`

- <a href="#VUID-vkCmdEncodeVideoKHR-constantQp-08274"
  id="VUID-vkCmdEncodeVideoKHR-constantQp-08274"></a>
  VUID-vkCmdEncodeVideoKHR-constantQp-08274  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`
  does not include
  `VK_VIDEO_ENCODE_H265_CAPABILITY_PER_SLICE_SEGMENT_CONSTANT_QP_BIT_KHR`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)::`constantQp`
  **must** have the same value for each element of the
  `pNaluSliceSegmentEntries` member of the
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo`

- <a href="#VUID-vkCmdEncodeVideoKHR-naluSliceSegmentEntryCount-08307"
  id="VUID-vkCmdEncodeVideoKHR-naluSliceSegmentEntryCount-08307"></a>
  VUID-vkCmdEncodeVideoKHR-naluSliceSegmentEntryCount-08307  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the
  `naluSliceSegmentEntryCount` member of the
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo` **must** be
  less than or equal to `minCodingBlockExtent.width` multiplied by
  `minCodingBlockExtent.height`

- <a href="#VUID-vkCmdEncodeVideoKHR-naluSliceSegmentEntryCount-08313"
  id="VUID-vkCmdEncodeVideoKHR-naluSliceSegmentEntryCount-08313"></a>
  VUID-vkCmdEncodeVideoKHR-naluSliceSegmentEntryCount-08313  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`
  does not include
  `VK_VIDEO_ENCODE_H265_CAPABILITY_ROW_UNALIGNED_SLICE_SEGMENT_BIT_KHR`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  the `naluSliceSegmentEntryCount` member of the
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
  structure included in the `pNext` chain of `pEncodeInfo` **must** be
  less than or equal to `minCodingBlockExtent.height`

- <a href="#VUID-vkCmdEncodeVideoKHR-pNext-08354"
  id="VUID-vkCmdEncodeVideoKHR-pNext-08354"></a>
  VUID-vkCmdEncodeVideoKHR-pNext-08354  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, the `pNext` chain of
  `pEncodeInfo` includes a
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
  structure, and `pEncodeInfo->referenceSlotCount` is greater than zero,
  then
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)::`pStdPictureInfo->pRefLists`
  **must** not be `NULL`

- <a href="#VUID-vkCmdEncodeVideoKHR-pNext-08344"
  id="VUID-vkCmdEncodeVideoKHR-pNext-08344"></a>
  VUID-vkCmdEncodeVideoKHR-pNext-08344  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, the `pNext` chain of
  `pEncodeInfo` includes a
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
  structure, and
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)::`pStdPictureInfo->pRefLists`
  is not `NULL`, then each element of the `RefPicList0` and
  `RefPicList1` array members of the
  `StdVideoEncodeH265ReferenceListsInfo` structure pointed to by
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)::`pStdPictureInfo->pRefLists`
  **must** either be `STD_VIDEO_H265_NO_REFERENCE_PICTURE` or **must**
  equal the `slotIndex` member of one of the elements of
  `pEncodeInfo->pReferenceSlots`

- <a href="#VUID-vkCmdEncodeVideoKHR-pNext-08355"
  id="VUID-vkCmdEncodeVideoKHR-pNext-08355"></a>
  VUID-vkCmdEncodeVideoKHR-pNext-08355  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, the `pNext` chain of
  `pEncodeInfo` includes a
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
  structure, and `pEncodeInfo->referenceSlotCount` is greater than zero,
  then the `slotIndex` member of each element of
  `pEncodeInfo->pReferenceSlots` **must** equal one of the elements of
  the `RefPicList0` or `RefPicList1` array members of the
  `StdVideoEncodeH265ReferenceListsInfo` structure pointed to by
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)::`pStdPictureInfo->pRefLists`

- <a href="#VUID-vkCmdEncodeVideoKHR-maxPPictureL0ReferenceCount-08345"
  id="VUID-vkCmdEncodeVideoKHR-maxPPictureL0ReferenceCount-08345"></a>
  VUID-vkCmdEncodeVideoKHR-maxPPictureL0ReferenceCount-08345  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxPPictureL0ReferenceCount`
  is zero, as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  `h265PictureType` and each element of `h265L0PictureTypes` and
  `h265L1PictureTypes` **must** not be `STD_VIDEO_H265_PICTURE_TYPE_P`

- <a href="#VUID-vkCmdEncodeVideoKHR-maxBPictureL0ReferenceCount-08346"
  id="VUID-vkCmdEncodeVideoKHR-maxBPictureL0ReferenceCount-08346"></a>
  VUID-vkCmdEncodeVideoKHR-maxBPictureL0ReferenceCount-08346  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxBPictureL0ReferenceCount`
  and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxL1ReferenceCount`
  are both zero, as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  `h265PictureType` and each element of `h265L0PictureTypes` and
  `h265L1PictureTypes` **must** not be `STD_VIDEO_H265_PICTURE_TYPE_B`

- <a href="#VUID-vkCmdEncodeVideoKHR-flags-08347"
  id="VUID-vkCmdEncodeVideoKHR-flags-08347"></a>
  VUID-vkCmdEncodeVideoKHR-flags-08347  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`
  does not include
  `VK_VIDEO_ENCODE_H265_CAPABILITY_B_FRAME_IN_L0_LIST_BIT_KHR`, as
  returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  each element of `h265L0PictureTypes` **must** not be
  `STD_VIDEO_H264_PICTURE_TYPE_B`

- <a href="#VUID-vkCmdEncodeVideoKHR-flags-08348"
  id="VUID-vkCmdEncodeVideoKHR-flags-08348"></a>
  VUID-vkCmdEncodeVideoKHR-flags-08348  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`
  does not include
  `VK_VIDEO_ENCODE_H265_CAPABILITY_B_FRAME_IN_L1_LIST_BIT_KHR`, as
  returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  each element of `h265L1PictureTypes` **must** not be
  `STD_VIDEO_H265_PICTURE_TYPE_B`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdEncodeVideoKHR-commandBuffer-parameter"
  id="VUID-vkCmdEncodeVideoKHR-commandBuffer-parameter"></a>
  VUID-vkCmdEncodeVideoKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdEncodeVideoKHR-pEncodeInfo-parameter"
  id="VUID-vkCmdEncodeVideoKHR-pEncodeInfo-parameter"></a>
  VUID-vkCmdEncodeVideoKHR-pEncodeInfo-parameter  
  `pEncodeInfo` **must** be a valid pointer to a valid
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html) structure

- <a href="#VUID-vkCmdEncodeVideoKHR-commandBuffer-recording"
  id="VUID-vkCmdEncodeVideoKHR-commandBuffer-recording"></a>
  VUID-vkCmdEncodeVideoKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdEncodeVideoKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdEncodeVideoKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdEncodeVideoKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support encode operations

- <a href="#VUID-vkCmdEncodeVideoKHR-renderpass"
  id="VUID-vkCmdEncodeVideoKHR-renderpass"></a>
  VUID-vkCmdEncodeVideoKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdEncodeVideoKHR-videocoding"
  id="VUID-vkCmdEncodeVideoKHR-videocoding"></a>
  VUID-vkCmdEncodeVideoKHR-videocoding  
  This command **must** only be called inside of a video coding scope

- <a href="#VUID-vkCmdEncodeVideoKHR-bufferlevel"
  id="VUID-vkCmdEncodeVideoKHR-bufferlevel"></a>
  VUID-vkCmdEncodeVideoKHR-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

| [Command Buffer Levels](#VkCommandBufferLevel) | [Render Pass Scope](#vkCmdBeginRenderPass) | [Video Coding Scope](#vkCmdBeginVideoCodingKHR) | [Supported Queue Types](#VkQueueFlagBits) | [Command Type](#fundamentals-queueoperation-command-types) |
|----|----|----|----|----|
| Primary | Outside | Inside | Encode | Action |

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdEncodeVideoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
