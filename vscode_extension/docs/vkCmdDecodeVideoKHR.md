# vkCmdDecodeVideoKHR(3) Manual Page

## Name

vkCmdDecodeVideoKHR - Launch a video decode operation



## <a href="#_c_specification" class="anchor"></a>C Specification

To launch video decode operations, call:

``` c
// Provided by VK_KHR_video_decode_queue
void vkCmdDecodeVideoKHR(
    VkCommandBuffer                             commandBuffer,
    const VkVideoDecodeInfoKHR*                 pDecodeInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to record the command.

- `pDecodeInfo` is a pointer to a
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html) structure specifying
  the parameters of the video decode operations.

## <a href="#_description" class="anchor"></a>Description

Each call issues one or more video decode operations. The implicit
parameter `opCount` corresponds to the number of video decode operations
issued by the command. After calling this command, the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active-query-index"
target="_blank" rel="noopener">active query index</a> of each <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
target="_blank" rel="noopener">active</a> query is incremented by
`opCount`.

Currently each call to this command results in the issue of a single
video decode operation.

If the bound video session was created with
`VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR` and the `pNext` chain
of `pDecodeInfo` includes a
[VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure
with its `queryPool` member specifying a valid `VkQueryPool` handle,
then this command will execute a query for each video decode operation
issued by it.

Active Reference Picture Information  
The list of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#active-reference-pictures"
target="_blank" rel="noopener">active reference pictures</a> used by a
video decode operation is a list of image subregions used as the source
of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-picture"
target="_blank" rel="noopener">reference picture</a> data and related
parameters, and is derived from the
[VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
structures provided as the elements of the
`pDecodeInfo->pReferenceSlots` array. For each element of
`pDecodeInfo->pReferenceSlots`, one or more elements are added to the
active reference picture list, as defined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-codec-specific-semantics"
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
decode operation is derived from the
[VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
structure pointed to by `pDecodeInfo->pSetupReferenceSlot`, if not
`NULL`, as defined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-codec-specific-semantics"
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
structure in `pDecodeInfo->pSetupReferenceSlot` is always required,
unless the video session was created with
[VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlot`
equal to zero. However, the DPB slot identified by
`pDecodeInfo->pSetupReferenceSlot->slotIndex` is only <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">activated</a> with the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reconstructed-picture"
target="_blank" rel="noopener">reconstructed picture</a> specified in
`pDecodeInfo->pSetupReferenceSlot->pPictureResource` if reference
picture setup is requested according to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-codec-specific-semantics"
target="_blank" rel="noopener">codec-specific semantics</a>.

If reconstructed picture information is specified, and
`pDecodeInfo->pSetupReferenceSlot->pPictureResource` refers to a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
target="_blank" rel="noopener">video picture resource</a> different than
that of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-output-picture"
target="_blank" rel="noopener">decode output picture</a>, but reference
picture setup is not requested, the contents of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
target="_blank" rel="noopener">video picture resource</a> corresponding
to the reconstructed picture will be undefined after the video decode
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
use it as temporary storage during the video decode operation even when
the reconstructed picture is not marked for future reference.</p></td>
</tr>
</tbody>
</table>

Decode Output Picture Information  
Information related to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-output-picture"
target="_blank" rel="noopener">decode output picture</a> used by a video
decode operation is derived from `pDecodeInfo->dstPictureResource` and
any codec-specific parameters provided in the `pDecodeInfo->pNext`
chain, as defined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-codec-specific-semantics"
target="_blank" rel="noopener">codec-specific semantics</a>, and
consists of the following:

- The image subregion within the image subresource <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">referred</a> to by the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
  target="_blank" rel="noopener">video picture resource</a> used as the
  decode output picture.

- The codec-specific picture information related to the decode output
  picture.

Several limiting values are defined below that are referenced by the
relevant valid usage statements of this command.

- Let `uint32_t activeReferencePictureCount` be the size of the list of
  active reference pictures used by the video decode operation. Unless
  otherwise defined, `activeReferencePictureCount` is set to the value
  of `pDecodeInfo->referenceSlotCount`.

  - If the bound video session was created with an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-profile"
    target="_blank" rel="noopener">H.264 decode profile</a>, then let
    `activeReferencePictureCount` be the value of
    `pDecodeInfo->referenceSlotCount` plus the number of elements of the
    `pDecodeInfo->pReferenceSlots` array that have a
    [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html)
    structure included in their `pNext` chain with both
    `pStdReferenceInfo->flags.top_field_flag` and
    `pStdReferenceInfo->flags.bottom_field_flag` set.

    <table>
    <colgroup>
    <col style="width: 50%" />
    <col style="width: 50%" />
    </colgroup>
    <tbody>
    <tr>
    <td class="icon"><em></em></td>
    <td class="content">Note
    <p>This means that the elements of
    <code>pDecodeInfo-&gt;pReferenceSlots</code> that include both a top and
    bottom field reference are counted as two separate active reference
    pictures, as described in the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-active-reference-picture-info"
    target="_blank" rel="noopener">active reference picture list
    construction rules for H.264 decode operations</a>.</p></td>
    </tr>
    </tbody>
    </table>

- Let `VkOffset2D codedOffsetGranularity` be the minimum alignment
  requirement for the coded offset of video picture resources. Unless
  otherwise defined, the value of the `x` and `y` members of
  `codedOffsetGranularity` are `0`.

  - If the bound video session was created with an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-profile"
    target="_blank" rel="noopener">H.264 decode profile</a> with a
    [VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264ProfileInfoKHR.html)::`pictureLayout`
    of
    `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_SEPARATE_PLANES_BIT_KHR`,
    then `codedOffsetGranularity` is equal to
    [VkVideoDecodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264CapabilitiesKHR.html)::`fieldOffsetGranularity`,
    as returned by
    [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
    for that video profile.

- Let `uint32_t dpbFrameUseCount[]` be an array of size `maxDpbSlots`,
  where `maxDpbSlots` is the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlots`
  the bound video session was created with, with each element indicating
  the number of times a frame associated with the corresponding DPB slot
  index is referred to by the video coding operation. Let the initial
  value of each element of the array be `0`.

  - If `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, then
    `dpbFrameUseCount[i]` is incremented by one, where `i` equals
    `pDecodeInfo->pSetupReferenceSlot->slotIndex`. If the bound video
    session object was created with an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-profile"
    target="_blank" rel="noopener">H.264 decode profile</a>, then
    `dpbFrameUseCount[i]` is decremented by one if either
    `pStdReferenceInfo->flags.top_field_flag` or
    `pStdReferenceInfo->flags.bottom_field_flag` is set in the
    [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html)
    structure in the `pDecodeInfo->pSetupReferenceSlot->pNext` chain.

  - For each element of `pDecodeInfo->pReferenceSlots`,
    `dpbFrameUseCount[i]` is incremented by one, where `i` equals the
    `slotIndex` member of the corresponding element. If the bound video
    session object was created with an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-profile"
    target="_blank" rel="noopener">H.264 decode profile</a>, then
    `dpbFrameUseCount[i]` is decremented by one if either
    `pStdReferenceInfo->flags.top_field_flag` or
    `pStdReferenceInfo->flags.bottom_field_flag` is set in the
    [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html)
    structure in the `pNext` chain of the corresponding element of
    `pDecodeInfo->pReferenceSlots`.

- Let `uint32_t dpbTopFieldUseCount[]` and
  `uint32_t dpbBottomFieldUseCount[]` be arrays of size `maxDpbSlots`,
  where `maxDpbSlots` is the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlots`
  the bound video session was created with, with each element indicating
  the number of times the top field or the bottom field, respectively,
  associated with the corresponding DPB slot index is referred to by the
  video coding operation. Let the initial value of each element of the
  arrays be `0`.

  - If the bound video session object was created with an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-profile"
    target="_blank" rel="noopener">H.264 decode profile</a> and
    `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, then perform the
    following:

    - If `pStdReferenceInfo->flags.top_field_flag` is set in the
      [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html)
      structure in the `pDecodeInfo->pSetupReferenceSlot->pNext` chain,
      then `dpbTopFieldUseCount[i]` is incremented by one, where `i`
      equals `pDecodeInfo->pSetupReferenceSlot->slotIndex`.

    - If `pStdReferenceInfo->flags.bottom_field_flag` is set in the
      [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html)
      structure in the `pDecodeInfo->pSetupReferenceSlot->pNext` chain,
      then `dpbBottomFieldUseCount[i]` is incremented by one, where `i`
      equals `pDecodeInfo->pSetupReferenceSlot->slotIndex`.

  - If the bound video session object was created with an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-profile"
    target="_blank" rel="noopener">H.264 decode profile</a>, then
    perform the following for each element of
    `pDecodeInfo->pReferenceSlots`:

    - If `pStdReferenceInfo->flags.top_field_flag` is set in the
      [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html)
      structure in the `pNext` chain of the element, then
      `dpbTopFieldUseCount[i]` is incremented by one, where `i` equals
      the `slotIndex` member of the element.

    - If `pStdReferenceInfo->flags.bottom_field_flag` is set in the
      [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html)
      structure in the `pNext` chain of the element, then
      `dpbBottomFieldUseCount[i]` is incremented by one, where `i`
      equals the `slotIndex` member of the element.

Valid Usage

- <a href="#VUID-vkCmdDecodeVideoKHR-None-08249"
  id="VUID-vkCmdDecodeVideoKHR-None-08249"></a>
  VUID-vkCmdDecodeVideoKHR-None-08249  
  The bound video session **must** have been created with a decode
  operation

- <a href="#VUID-vkCmdDecodeVideoKHR-None-07011"
  id="VUID-vkCmdDecodeVideoKHR-None-07011"></a>
  VUID-vkCmdDecodeVideoKHR-None-07011  
  The bound video session **must** not be in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-uninitialized"
  target="_blank" rel="noopener">uninitialized</a> state at the time the
  command is executed on the device

- <a href="#VUID-vkCmdDecodeVideoKHR-opCount-07134"
  id="VUID-vkCmdDecodeVideoKHR-opCount-07134"></a>
  VUID-vkCmdDecodeVideoKHR-opCount-07134  
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

- <a href="#VUID-vkCmdDecodeVideoKHR-pNext-08365"
  id="VUID-vkCmdDecodeVideoKHR-pNext-08365"></a>
  VUID-vkCmdDecodeVideoKHR-pNext-08365  
  If the bound video session was created with
  `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`, and the `pNext`
  chain of `pDecodeInfo` includes a
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure
  with its `queryPool` member specifying a valid `VkQueryPool` handle,
  then
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html)::queryCount
  **must** equal `opCount`

- <a href="#VUID-vkCmdDecodeVideoKHR-pNext-08366"
  id="VUID-vkCmdDecodeVideoKHR-pNext-08366"></a>
  VUID-vkCmdDecodeVideoKHR-pNext-08366  
  If the bound video session was created with
  `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`, and the `pNext`
  chain of `pDecodeInfo` includes a
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure
  with its `queryPool` member specifying a valid `VkQueryPool` handle,
  then all the queries used by the command, as specified by the
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure,
  **must** be *unavailable*

- <a href="#VUID-vkCmdDecodeVideoKHR-queryType-08367"
  id="VUID-vkCmdDecodeVideoKHR-queryType-08367"></a>
  VUID-vkCmdDecodeVideoKHR-queryType-08367  
  If the bound video session was created with
  `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`, then the `queryType`
  used to create the `queryPool` specified in the
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure
  included in the `pNext` chain of `pDecodeInfo` **must** be
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`

- <a href="#VUID-vkCmdDecodeVideoKHR-queryPool-08368"
  id="VUID-vkCmdDecodeVideoKHR-queryPool-08368"></a>
  VUID-vkCmdDecodeVideoKHR-queryPool-08368  
  If the bound video session was created with
  `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`, then the `queryPool`
  specified in the
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure
  included in the `pNext` chain of `pDecodeInfo` **must** have been
  created with a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)
  structure included in the `pNext` chain of
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html) identical to the
  one specified in
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`pVideoProfile`
  the bound video session was created with

- <a href="#VUID-vkCmdDecodeVideoKHR-queryType-08369"
  id="VUID-vkCmdDecodeVideoKHR-queryType-08369"></a>
  VUID-vkCmdDecodeVideoKHR-queryType-08369  
  If the bound video session was created with
  `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`, and the `queryType`
  used to create the `queryPool` specified in the
  [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html) structure
  included in the `pNext` chain of `pDecodeInfo` is
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`, then the `VkCommandPool` that
  `commandBuffer` was allocated from **must** have been created with a
  queue family index that supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-result-status-only"
  target="_blank" rel="noopener">result status queries</a>, as indicated
  by
  [VkQueueFamilyQueryResultStatusPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyQueryResultStatusPropertiesKHR.html)::`queryResultStatusSupport`

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07135"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07135"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07135  
  `pDecodeInfo->srcBuffer` **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profile-compatibility"
  target="_blank" rel="noopener">compatible</a> with the video profile
  the bound video session was created with

- <a href="#VUID-vkCmdDecodeVideoKHR-commandBuffer-07136"
  id="VUID-vkCmdDecodeVideoKHR-commandBuffer-07136"></a>
  VUID-vkCmdDecodeVideoKHR-commandBuffer-07136  
  If `commandBuffer` is an unprotected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, then `pDecodeInfo->srcBuffer` **must** not be a
  protected buffer

- <a href="#VUID-vkCmdDecodeVideoKHR-commandBuffer-07137"
  id="VUID-vkCmdDecodeVideoKHR-commandBuffer-07137"></a>
  VUID-vkCmdDecodeVideoKHR-commandBuffer-07137  
  If `commandBuffer` is a protected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, then `pDecodeInfo->srcBuffer` **must** be a protected
  buffer

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07138"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07138"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07138  
  `pDecodeInfo->srcBufferOffset` **must** be an integer multiple of
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`minBitstreamBufferOffsetAlignment`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07139"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07139"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07139  
  `pDecodeInfo->srcBufferRange` **must** be an integer multiple of
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`minBitstreamBufferSizeAlignment`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07140"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07140"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07140  
  If `pDecodeInfo->pSetupReferenceSlot` is not `NULL` and
  [VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeCapabilitiesKHR.html)::`flags`
  does not include
  `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR`, as
  returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the bound video session was created with, then
  the video picture resources specified by
  `pDecodeInfo->dstPictureResource` and
  `pDecodeInfo->pSetupReferenceSlot->pPictureResource` **must** not <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-matching"
  target="_blank" rel="noopener">match</a>

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07141"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07141"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07141  
  If `pDecodeInfo->pSetupReferenceSlot` is not `NULL` and none of the
  following is true:

  - [VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeCapabilitiesKHR.html)::`flags`
    includes
    `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR`, as
    returned by
    [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
    for the video profile the bound video session was created with

  - the bound video session was created with the video codec operation
    `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR` and
    [VkVideoDecodeAV1ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1ProfileInfoKHR.html)::`filmGrainSupport`
    set to `VK_TRUE`, and <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-film-grain"
    target="_blank" rel="noopener">film grain</a> is enabled for the
    decoded picture

  then the video picture resources specified by
  `pDecodeInfo->dstPictureResource` and
  `pDecodeInfo->pSetupReferenceSlot->pPictureResource` **must** <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-matching"
  target="_blank" rel="noopener">match</a>

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07142"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07142"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07142  
  `pDecodeInfo->dstPictureResource.imageViewBinding` **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profile-compatibility"
  target="_blank" rel="noopener">compatible</a> with the video profile
  the bound video session was created with

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07143"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07143"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07143  
  The format of `pDecodeInfo->dstPictureResource.imageViewBinding`
  **must** match the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`pictureFormat`
  the bound video session was created with

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07144"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07144"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07144  
  `pDecodeInfo->dstPictureResource.codedOffset` **must** be an integer
  multiple of `codedOffsetGranularity`

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07145"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07145"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07145  
  `pDecodeInfo->dstPictureResource.codedExtent` **must** be between
  `minCodedExtent` and `maxCodedExtent`, inclusive, the bound video
  session was created with

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07146"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07146"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07146  
  `pDecodeInfo->dstPictureResource.imageViewBinding` **must** have been
  created with `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`

- <a href="#VUID-vkCmdDecodeVideoKHR-commandBuffer-07147"
  id="VUID-vkCmdDecodeVideoKHR-commandBuffer-07147"></a>
  VUID-vkCmdDecodeVideoKHR-commandBuffer-07147  
  If `commandBuffer` is an unprotected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, then `pDecodeInfo->dstPictureResource.imageViewBinding`
  **must** not have been created from a protected image

- <a href="#VUID-vkCmdDecodeVideoKHR-commandBuffer-07148"
  id="VUID-vkCmdDecodeVideoKHR-commandBuffer-07148"></a>
  VUID-vkCmdDecodeVideoKHR-commandBuffer-07148  
  If `commandBuffer` is a protected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, then `pDecodeInfo->dstPictureResource.imageViewBinding`
  **must** have been created from a protected image

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-08376"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-08376"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-08376  
  `pDecodeInfo->pSetupReferenceSlot` **must** not be `NULL` unless the
  bound video session was created with
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlots`
  equal to zero

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07170"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07170"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07170  
  If `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, then
  `pDecodeInfo->pSetupReferenceSlot->slotIndex` **must** be less than
  the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlots`
  specified when the bound video session was created

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07173"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07173"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07173  
  If `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, then
  `pDecodeInfo->pSetupReferenceSlot->pPictureResource->codedOffset`
  **must** be an integer multiple of `codedOffsetGranularity`

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07149"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07149"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07149  
  If `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, then
  `pDecodeInfo->pSetupReferenceSlot->pPictureResource` **must** <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-matching"
  target="_blank" rel="noopener">match</a> one of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#bound-reference-picture-resources"
  target="_blank" rel="noopener">bound reference picture resource</a>

- <a href="#VUID-vkCmdDecodeVideoKHR-activeReferencePictureCount-07150"
  id="VUID-vkCmdDecodeVideoKHR-activeReferencePictureCount-07150"></a>
  VUID-vkCmdDecodeVideoKHR-activeReferencePictureCount-07150  
  `activeReferencePictureCount` **must** be less than or equal to the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxActiveReferencePictures`
  specified when the bound video session was created

- <a href="#VUID-vkCmdDecodeVideoKHR-slotIndex-07256"
  id="VUID-vkCmdDecodeVideoKHR-slotIndex-07256"></a>
  VUID-vkCmdDecodeVideoKHR-slotIndex-07256  
  The `slotIndex` member of each element of
  `pDecodeInfo->pReferenceSlots` **must** be less than the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlots`
  specified when the bound video session was created

- <a href="#VUID-vkCmdDecodeVideoKHR-codedOffset-07257"
  id="VUID-vkCmdDecodeVideoKHR-codedOffset-07257"></a>
  VUID-vkCmdDecodeVideoKHR-codedOffset-07257  
  The `codedOffset` member of the
  [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html)
  structure pointed to by the `pPictureResource` member of each element
  of `pDecodeInfo->pReferenceSlots` **must** be an integer multiple of
  `codedOffsetGranularity`

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07151"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07151"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07151  
  The `pPictureResource` member of each element of
  `pDecodeInfo->pReferenceSlots` **must** <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-matching"
  target="_blank" rel="noopener">match</a> one of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#bound-reference-picture-resources"
  target="_blank" rel="noopener">bound reference picture resource</a>
  associated with the DPB slot index specified in the `slotIndex` member
  of that element

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07264"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07264"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07264  
  Each video picture resource corresponding to the `pPictureResource`
  member specified in the elements of `pDecodeInfo->pReferenceSlots`
  **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-uniqueness"
  target="_blank" rel="noopener">unique</a> within
  `pDecodeInfo->pReferenceSlots`

- <a href="#VUID-vkCmdDecodeVideoKHR-dpbFrameUseCount-07176"
  id="VUID-vkCmdDecodeVideoKHR-dpbFrameUseCount-07176"></a>
  VUID-vkCmdDecodeVideoKHR-dpbFrameUseCount-07176  
  All elements of `dpbFrameUseCount` **must** be less than or equal to
  `1`

- <a href="#VUID-vkCmdDecodeVideoKHR-dpbTopFieldUseCount-07177"
  id="VUID-vkCmdDecodeVideoKHR-dpbTopFieldUseCount-07177"></a>
  VUID-vkCmdDecodeVideoKHR-dpbTopFieldUseCount-07177  
  All elements of `dpbTopFieldUseCount` **must** be less than or equal
  to `1`

- <a href="#VUID-vkCmdDecodeVideoKHR-dpbBottomFieldUseCount-07178"
  id="VUID-vkCmdDecodeVideoKHR-dpbBottomFieldUseCount-07178"></a>
  VUID-vkCmdDecodeVideoKHR-dpbBottomFieldUseCount-07178  
  All elements of `dpbBottomFieldUseCount` **must** be less than or
  equal to `1`

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07252"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07252"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07252  
  If `pDecodeInfo->pSetupReferenceSlot` is `NULL` or
  `pDecodeInfo->pSetupReferenceSlot->pPictureResource` does not <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">refer</a> to the same image subresource
  as `pDecodeInfo->dstPictureResource`, then the image subresource <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">referred</a> to by
  `pDecodeInfo->dstPictureResource` **must** be in the
  `VK_IMAGE_LAYOUT_VIDEO_DECODE_DST_KHR` layout at the time the video
  decode operation is executed on the device

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07253"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07253"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07253  
  If `pDecodeInfo->pSetupReferenceSlot` is not `NULL` and
  `pDecodeInfo->pSetupReferenceSlot->pPictureResource` <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">refers</a> to the same image
  subresource as `pDecodeInfo->dstPictureResource`, then the image
  subresource <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">referred</a> to by
  `pDecodeInfo->dstPictureResource` **must** be in the
  `VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR` layout at the time the video
  decode operation is executed on the device

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07254"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07254"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07254  
  If `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, then the image
  subresource <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">referred</a> to by
  `pDecodeInfo->pSetupReferenceSlot->pPictureResource` **must** be in
  the `VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR` layout at the time the
  video decode operation is executed on the device

- <a href="#VUID-vkCmdDecodeVideoKHR-pPictureResource-07255"
  id="VUID-vkCmdDecodeVideoKHR-pPictureResource-07255"></a>
  VUID-vkCmdDecodeVideoKHR-pPictureResource-07255  
  The image subresource <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-image-subresource-reference"
  target="_blank" rel="noopener">referred</a> to by the
  `pPictureResource` member of each element of
  `pDecodeInfo->pReferenceSlots` **must** be in the
  `VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR` layout at the time the video
  decode operation is executed on the device

- <a href="#VUID-vkCmdDecodeVideoKHR-pNext-07152"
  id="VUID-vkCmdDecodeVideoKHR-pNext-07152"></a>
  VUID-vkCmdDecodeVideoKHR-pNext-07152  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the `pNext` chain
  of `pDecodeInfo` **must** include a
  [VkVideoDecodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264PictureInfoKHR.html)
  structure

- <a href="#VUID-vkCmdDecodeVideoKHR-None-07258"
  id="VUID-vkCmdDecodeVideoKHR-None-07258"></a>
  VUID-vkCmdDecodeVideoKHR-None-07258  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` but was not created
  with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-interlaced-support"
  target="_blank" rel="noopener">interlaced frame support</a>, then the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-output-picture-info"
  target="_blank" rel="noopener">decode output picture</a> **must**
  represent a frame

- <a href="#VUID-vkCmdDecodeVideoKHR-pSliceOffsets-07153"
  id="VUID-vkCmdDecodeVideoKHR-pSliceOffsets-07153"></a>
  VUID-vkCmdDecodeVideoKHR-pSliceOffsets-07153  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then all elements of
  the `pSliceOffsets` member of the
  [VkVideoDecodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264PictureInfoKHR.html)
  structure included in the `pNext` chain of `pDecodeInfo` **must** be
  less than `pDecodeInfo->srcBufferRange`

- <a
  href="#VUID-vkCmdDecodeVideoKHR-StdVideoH264SequenceParameterSet-07154"
  id="VUID-vkCmdDecodeVideoKHR-StdVideoH264SequenceParameterSet-07154"></a>
  VUID-vkCmdDecodeVideoKHR-StdVideoH264SequenceParameterSet-07154  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the bound video
  session parameters object **must** contain a
  `StdVideoH264SequenceParameterSet` entry with `seq_parameter_set_id`
  matching `StdVideoDecodeH264PictureInfo`::`seq_parameter_set_id` that
  is provided in the `pStdPictureInfo` member of the
  [VkVideoDecodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264PictureInfoKHR.html)
  structure included in the `pNext` chain of `pDecodeInfo`

- <a
  href="#VUID-vkCmdDecodeVideoKHR-StdVideoH264PictureParameterSet-07155"
  id="VUID-vkCmdDecodeVideoKHR-StdVideoH264PictureParameterSet-07155"></a>
  VUID-vkCmdDecodeVideoKHR-StdVideoH264PictureParameterSet-07155  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the bound video
  session parameters object **must** contain a
  `StdVideoH264PictureParameterSet` entry with `seq_parameter_set_id`
  and `pic_parameter_set_id` matching
  `StdVideoDecodeH264PictureInfo`::`seq_parameter_set_id` and
  `StdVideoDecodeH264PictureInfo`::`pic_parameter_set_id`, respectively,
  that are provided in the `pStdPictureInfo` member of the
  [VkVideoDecodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264PictureInfoKHR.html)
  structure included in the `pNext` chain of `pDecodeInfo`

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07156"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07156"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07156  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` and
  `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, then the `pNext`
  chain of `pDecodeInfo->pSetupReferenceSlot` **must** include a
  [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html)
  structure

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07259"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07259"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07259  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` but was not created
  with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-interlaced-support"
  target="_blank" rel="noopener">interlaced frame support</a>, and
  `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-reconstructed-picture-info"
  target="_blank" rel="noopener">reconstructed picture</a> **must**
  represent a frame

- <a href="#VUID-vkCmdDecodeVideoKHR-pNext-07157"
  id="VUID-vkCmdDecodeVideoKHR-pNext-07157"></a>
  VUID-vkCmdDecodeVideoKHR-pNext-07157  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the `pNext` chain
  of each element of `pDecodeInfo->pReferenceSlots` **must** include a
  [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html)
  structure

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07260"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07260"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07260  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` but was not created
  with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-interlaced-support"
  target="_blank" rel="noopener">interlaced frame support</a>, then each
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-active-reference-picture-info"
  target="_blank" rel="noopener">active reference picture</a>
  corresponding to the elements of `pDecodeInfo->pReferenceSlots`
  **must** represent a frame

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07261"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07261"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07261  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`,
  `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-output-picture-info"
  target="_blank" rel="noopener">decode output picture</a> represents a
  frame, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-reconstructed-picture-info"
  target="_blank" rel="noopener">reconstructed picture</a> **must** also
  represent a frame

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07262"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07262"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07262  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`,
  `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-output-picture-info"
  target="_blank" rel="noopener">decode output picture</a> represents a
  top field, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-reconstructed-picture-info"
  target="_blank" rel="noopener">reconstructed picture</a> **must** also
  represent a top field

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07263"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07263"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07263  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`,
  `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-output-picture-info"
  target="_blank" rel="noopener">decode output picture</a> represents a
  bottom field, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-reconstructed-picture-info"
  target="_blank" rel="noopener">reconstructed picture</a> **must** also
  represent a bottom field

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07266"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07266"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07266  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` and an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-active-reference-picture-info"
  target="_blank" rel="noopener">active reference picture</a>
  corresponding to any element of `pDecodeInfo->pReferenceSlots`
  represents a frame, then the DPB slot index of the bound video session
  specified by the `slotIndex` member of that element **must** be
  currently associated with a frame picture <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-matching"
  target="_blank" rel="noopener">matching</a> the video picture resource
  specified by the `pPictureResource` member of the same element at the
  time the command is executed on the device

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07267"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07267"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07267  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` and an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-active-reference-picture-info"
  target="_blank" rel="noopener">active reference picture</a>
  corresponding to any element of `pDecodeInfo->pReferenceSlots`
  represents a top field, then the DPB slot index of the bound video
  session specified by the `slotIndex` member of that element **must**
  be currently associated with a top field picture <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-matching"
  target="_blank" rel="noopener">matching</a> the video picture resource
  specified by the `pPictureResource` member of the same element at the
  time the command is executed on the device

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07268"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07268"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07268  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` and an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-active-reference-picture-info"
  target="_blank" rel="noopener">active reference picture</a>
  corresponding to any element of `pDecodeInfo->pReferenceSlots`
  represents a bottom field, then the DPB slot index of the bound video
  session specified by the `slotIndex` member of that element **must**
  be currently associated with a bottom field picture <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-matching"
  target="_blank" rel="noopener">matching</a> the video picture resource
  specified by the `pPictureResource` member of the same element at the
  time the command is executed on the device

- <a href="#VUID-vkCmdDecodeVideoKHR-pNext-07158"
  id="VUID-vkCmdDecodeVideoKHR-pNext-07158"></a>
  VUID-vkCmdDecodeVideoKHR-pNext-07158  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the `pNext` chain
  of `pDecodeInfo` **must** include a
  [VkVideoDecodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265PictureInfoKHR.html)
  structure

- <a href="#VUID-vkCmdDecodeVideoKHR-pSliceSegmentOffsets-07159"
  id="VUID-vkCmdDecodeVideoKHR-pSliceSegmentOffsets-07159"></a>
  VUID-vkCmdDecodeVideoKHR-pSliceSegmentOffsets-07159  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then all elements of
  the `pSliceSegmentOffsets` member of the
  [VkVideoDecodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265PictureInfoKHR.html)
  structure included in the `pNext` chain of `pDecodeInfo` **must** be
  less than `pDecodeInfo->srcBufferRange`

- <a href="#VUID-vkCmdDecodeVideoKHR-StdVideoH265VideoParameterSet-07160"
  id="VUID-vkCmdDecodeVideoKHR-StdVideoH265VideoParameterSet-07160"></a>
  VUID-vkCmdDecodeVideoKHR-StdVideoH265VideoParameterSet-07160  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the bound video
  session parameters object **must** contain a
  `StdVideoH265VideoParameterSet` entry with
  `vps_video_parameter_set_id` matching
  `StdVideoDecodeH265PictureInfo`::`sps_video_parameter_set_id` that is
  provided in the `pStdPictureInfo` member of the
  [VkVideoDecodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265PictureInfoKHR.html)
  structure included in the `pNext` chain of `pDecodeInfo`

- <a
  href="#VUID-vkCmdDecodeVideoKHR-StdVideoH265SequenceParameterSet-07161"
  id="VUID-vkCmdDecodeVideoKHR-StdVideoH265SequenceParameterSet-07161"></a>
  VUID-vkCmdDecodeVideoKHR-StdVideoH265SequenceParameterSet-07161  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the bound video
  session parameters object **must** contain a
  `StdVideoH265SequenceParameterSet` entry with
  `sps_video_parameter_set_id` and `sps_seq_parameter_set_id` matching
  `StdVideoDecodeH265PictureInfo`::`sps_video_parameter_set_id` and
  `StdVideoDecodeH265PictureInfo`::`pps_seq_parameter_set_id`,
  respectively, that are provided in the `pStdPictureInfo` member of the
  [VkVideoDecodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265PictureInfoKHR.html)
  structure included in the `pNext` chain of `pDecodeInfo`

- <a
  href="#VUID-vkCmdDecodeVideoKHR-StdVideoH265PictureParameterSet-07162"
  id="VUID-vkCmdDecodeVideoKHR-StdVideoH265PictureParameterSet-07162"></a>
  VUID-vkCmdDecodeVideoKHR-StdVideoH265PictureParameterSet-07162  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the bound video
  session parameters object **must** contain a
  `StdVideoH265PictureParameterSet` entry with
  `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and
  `pps_pic_parameter_set_id` matching
  `StdVideoDecodeH265PictureInfo`::`sps_video_parameter_set_id`,
  `StdVideoDecodeH265PictureInfo`::`pps_seq_parameter_set_id`, and
  `StdVideoDecodeH265PictureInfo`::`pps_pic_parameter_set_id`,
  respectively, that are provided in the `pStdPictureInfo` member of the
  [VkVideoDecodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265PictureInfoKHR.html)
  structure included in the `pNext` chain of `pDecodeInfo`

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07163"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07163"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-07163  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR` and
  `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, then the `pNext`
  chain of `pDecodeInfo->pSetupReferenceSlot` **must** include a
  [VkVideoDecodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265DpbSlotInfoKHR.html)
  structure

- <a href="#VUID-vkCmdDecodeVideoKHR-pNext-07164"
  id="VUID-vkCmdDecodeVideoKHR-pNext-07164"></a>
  VUID-vkCmdDecodeVideoKHR-pNext-07164  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the `pNext` chain
  of each element of `pDecodeInfo->pReferenceSlots` **must** include a
  [VkVideoDecodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265DpbSlotInfoKHR.html)
  structure

- <a href="#VUID-vkCmdDecodeVideoKHR-filmGrainSupport-09248"
  id="VUID-vkCmdDecodeVideoKHR-filmGrainSupport-09248"></a>
  VUID-vkCmdDecodeVideoKHR-filmGrainSupport-09248  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR` and
  [VkVideoDecodeAV1ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1ProfileInfoKHR.html)::`filmGrainSupport`
  set to `VK_FALSE`, then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-film-grain"
  target="_blank" rel="noopener">film grain</a> **must** not be enabled
  for the decoded picture

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-09249"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-09249"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-09249  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`,
  `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-film-grain"
  target="_blank" rel="noopener">film grain</a> is enabled for the
  decoded picture, then the video picture resources specified by
  `pDecodeInfo->dstPictureResource` and
  `pDecodeInfo->pSetupReferenceSlot->pPictureResource` **must** not <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-matching"
  target="_blank" rel="noopener">match</a>

- <a href="#VUID-vkCmdDecodeVideoKHR-pNext-09250"
  id="VUID-vkCmdDecodeVideoKHR-pNext-09250"></a>
  VUID-vkCmdDecodeVideoKHR-pNext-09250  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then the `pNext` chain
  of `pDecodeInfo` **must** include a
  [VkVideoDecodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1PictureInfoKHR.html)
  structure

- <a href="#VUID-vkCmdDecodeVideoKHR-frameHeaderOffset-09251"
  id="VUID-vkCmdDecodeVideoKHR-frameHeaderOffset-09251"></a>
  VUID-vkCmdDecodeVideoKHR-frameHeaderOffset-09251  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then the
  `frameHeaderOffset` member of the
  [VkVideoDecodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1PictureInfoKHR.html)
  structure included in the `pNext` chain of `pDecodeInfo` **must** be
  less than the minimum of `pDecodeInfo->srcBufferRange`

- <a href="#VUID-vkCmdDecodeVideoKHR-pTileOffsets-09253"
  id="VUID-vkCmdDecodeVideoKHR-pTileOffsets-09253"></a>
  VUID-vkCmdDecodeVideoKHR-pTileOffsets-09253  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then all elements of
  the `pTileOffsets` member of the
  [VkVideoDecodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1PictureInfoKHR.html)
  structure included in the `pNext` chain of `pDecodeInfo` **must** be
  less than `pDecodeInfo->srcBufferRange`

- <a href="#VUID-vkCmdDecodeVideoKHR-pTileOffsets-09252"
  id="VUID-vkCmdDecodeVideoKHR-pTileOffsets-09252"></a>
  VUID-vkCmdDecodeVideoKHR-pTileOffsets-09252  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then for each element i
  of the `pTileOffsets` and `pTileSizes` members of the
  [VkVideoDecodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1PictureInfoKHR.html)
  structure included in the `pNext` chain of `pDecodeInfo` the sum of
  `pTileOffsets`\[i\] and `pTileSizes`\[i\] **must** be less than or
  equal to `pDecodeInfo->srcBufferRange`

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-09254"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-09254"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-09254  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR` and
  `pDecodeInfo->pSetupReferenceSlot` is not `NULL`, then the `pNext`
  chain of `pDecodeInfo->pSetupReferenceSlot` **must** include a
  [VkVideoDecodeAV1DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1DpbSlotInfoKHR.html)
  structure

- <a href="#VUID-vkCmdDecodeVideoKHR-pNext-09255"
  id="VUID-vkCmdDecodeVideoKHR-pNext-09255"></a>
  VUID-vkCmdDecodeVideoKHR-pNext-09255  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then the `pNext` chain
  of each element of `pDecodeInfo->pReferenceSlots` **must** include a
  [VkVideoDecodeAV1DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1DpbSlotInfoKHR.html)
  structure

- <a href="#VUID-vkCmdDecodeVideoKHR-referenceNameSlotIndices-09262"
  id="VUID-vkCmdDecodeVideoKHR-referenceNameSlotIndices-09262"></a>
  VUID-vkCmdDecodeVideoKHR-referenceNameSlotIndices-09262  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then each element of
  the `referenceNameSlotIndices` array member of the
  [VkVideoDecodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1PictureInfoKHR.html)
  structure included in the `pNext` chain of `pDecodeInfo` **must**
  either be negative or **must** equal the `slotIndex` member of one of
  the elements of `pDecodeInfo->pReferenceSlots`

- <a href="#VUID-vkCmdDecodeVideoKHR-slotIndex-09263"
  id="VUID-vkCmdDecodeVideoKHR-slotIndex-09263"></a>
  VUID-vkCmdDecodeVideoKHR-slotIndex-09263  
  If the bound video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then the `slotIndex`
  member of each element of `pDecodeInfo->pReferenceSlots` **must**
  equal one of the elements of the `referenceNameSlotIndices` array
  member of the
  [VkVideoDecodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1PictureInfoKHR.html)
  structure included in the `pNext` chain of `pDecodeInfo`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdDecodeVideoKHR-commandBuffer-parameter"
  id="VUID-vkCmdDecodeVideoKHR-commandBuffer-parameter"></a>
  VUID-vkCmdDecodeVideoKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdDecodeVideoKHR-pDecodeInfo-parameter"
  id="VUID-vkCmdDecodeVideoKHR-pDecodeInfo-parameter"></a>
  VUID-vkCmdDecodeVideoKHR-pDecodeInfo-parameter  
  `pDecodeInfo` **must** be a valid pointer to a valid
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html) structure

- <a href="#VUID-vkCmdDecodeVideoKHR-commandBuffer-recording"
  id="VUID-vkCmdDecodeVideoKHR-commandBuffer-recording"></a>
  VUID-vkCmdDecodeVideoKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdDecodeVideoKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdDecodeVideoKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdDecodeVideoKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support decode operations

- <a href="#VUID-vkCmdDecodeVideoKHR-renderpass"
  id="VUID-vkCmdDecodeVideoKHR-renderpass"></a>
  VUID-vkCmdDecodeVideoKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdDecodeVideoKHR-videocoding"
  id="VUID-vkCmdDecodeVideoKHR-videocoding"></a>
  VUID-vkCmdDecodeVideoKHR-videocoding  
  This command **must** only be called inside of a video coding scope

- <a href="#VUID-vkCmdDecodeVideoKHR-bufferlevel"
  id="VUID-vkCmdDecodeVideoKHR-bufferlevel"></a>
  VUID-vkCmdDecodeVideoKHR-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

| [Command Buffer Levels](#VkCommandBufferLevel) | [Render Pass Scope](#vkCmdBeginRenderPass) | [Video Coding Scope](#vkCmdBeginVideoCodingKHR) | [Supported Queue Types](#VkQueueFlagBits) | [Command Type](#fundamentals-queueoperation-command-types) |
|----|----|----|----|----|
| Primary | Outside | Inside | Decode | Action |

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_queue.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdDecodeVideoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
