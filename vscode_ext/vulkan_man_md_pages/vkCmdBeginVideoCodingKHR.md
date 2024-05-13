# vkCmdBeginVideoCodingKHR(3) Manual Page

## Name

vkCmdBeginVideoCodingKHR - Begin video coding scope



## <a href="#_c_specification" class="anchor"></a>C Specification

To begin a video coding scope, call:

``` c
// Provided by VK_KHR_video_queue
void vkCmdBeginVideoCodingKHR(
    VkCommandBuffer                             commandBuffer,
    const VkVideoBeginCodingInfoKHR*            pBeginInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to record the command.

- `pBeginInfo` is a pointer to a
  [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html) structure
  specifying the parameters of the video coding scope, including the
  video session and video session parameters object to use.

## <a href="#_description" class="anchor"></a>Description

After beginning a video coding scope, the video session object specified
in `pBeginInfo->videoSession` is *bound* to the command buffer, and the
command buffer is ready to record video coding operations. Similarly, if
`pBeginInfo->videoSessionParameters` is not `VK_NULL_HANDLE`, it is also
*bound* to the command buffer, and video coding operations **can** refer
to the codec-specific parameters stored in it.

This command also establishes the set of *bound reference picture
resources* that **can** be used as <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reconstructed-picture"
target="_blank" rel="noopener">reconstructed pictures</a> or <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-picture"
target="_blank" rel="noopener">reference pictures</a> within the video
coding scope. Each element of this set consists of a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
target="_blank" rel="noopener">video picture resource</a> and the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
target="_blank" rel="noopener">DPB slot</a> index associated with it, if
there is one.

The set of bound reference picture resources is immutable within a video
coding scope, however, the DPB slot index associated with any of the
bound reference picture resources **can** change during the video coding
scope in response to video coding operations.

The [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
structures provided as the elements of `pBeginInfo->pReferenceSlots` are
interpreted by this command as follows:

- If `slotIndex` is non-negative and `pPictureResource` is not `NULL`,
  then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
  target="_blank" rel="noopener">video picture resource</a> defined by
  the
  [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html)
  structure pointed to by `pPictureResource` is added to the set of
  bound reference picture resources and is associated with the DPB slot
  index specified in `slotIndex`.

- If `slotIndex` is non-negative and `pPictureResource` is `NULL`, then
  the DPB slot with index `slotIndex` is <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
  target="_blank" rel="noopener">deactivated</a> by this command.

- If `slotIndex` is negative and `pPictureResource` is not `NULL`, then
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
  target="_blank" rel="noopener">video picture resource</a> defined by
  the
  [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html)
  structure pointed to by `pPictureResource` is added to the set of
  bound reference picture resources without an associated DPB slot. Such
  a picture resource **can** be subsequently used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reconstructed-picture"
  target="_blank" rel="noopener">reconstructed picture</a> to associate
  it with a DPB slot.

- If `slotIndex` is negative and `pPictureResource` is `NULL`, then the
  element is ignored.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>It is possible for multiple bound reference picture resources to be
associated with the same DPB slot index, or for a single bound reference
picture to refer to multiple separate reference pictures. For example,
in case of an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-profile"
target="_blank" rel="noopener">H.264 decode profile</a> with <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-interlaced-support"
target="_blank" rel="noopener">interlaced frame support</a> a single DPB
slot can refer to two separate pictures for the top and bottom fields.
Depending on the picture layout used by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-profile"
target="_blank" rel="noopener">H.264 decode profile</a>, the following
special cases <strong>may</strong> arise:</p>
<ul>
<li><p>If the picture layout is
<code>VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_INTERLEAVED_LINES_BIT_KHR</code>,
then the top and bottom field pictures are physically co-located in the
same video picture resource with even scanlines corresponding to the top
field and odd scanlines corresponding to the bottom field,
respectively.</p></li>
<li><p>If the picture layout is
<code>VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_SEPARATE_PLANES_BIT_KHR</code>,
then the top and bottom field pictures are stored in separate video
picture resources (in separate subregions of the same image layer, in
separate layers of the same image, or in entirely separate images),
hence two elements of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html">VkVideoBeginCodingInfoKHR</a>::<code>pReferenceSlots</code>
<strong>can</strong> contain the same <code>slotIndex</code> but specify
different video picture resources in their <code>pPictureResource</code>
members.</p></li>
</ul></td>
</tr>
</tbody>
</table>

All non-negative `slotIndex` values specified in the elements of
`pBeginInfo->pReferenceSlots` **must** identify DPB slots of the video
session that are in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">active state</a> at the time this command
is executed on the device.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The application does not have to specify an entry in
<code>pBeginInfo-&gt;pReferenceSlots</code> corresponding to all active
DPB slots of the video session, but only for those which are intended to
be used in the video coding scope. This way the application can avoid
any potential runtime cost associated with binding the corresponding
picture resources to the command buffer.</p></td>
</tr>
</tbody>
</table>

In case of a video encode session, the application is also responsible
for providing information about the current <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-state"
target="_blank" rel="noopener">rate control state</a> configured for the
video session by including an instance of the
[VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html)
structure in the `pNext` chain of `pBeginInfo`. If no
[VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html)
is included, then the presence of an empty
[VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html)
structure is implied which indicates that the current <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
target="_blank" rel="noopener">rate control mode</a> is
`VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DEFAULT_KHR`. The specified state
**must** <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-state-matching"
target="_blank" rel="noopener">match</a> the effective rate control
state configured for the video session at the time the recorded command
is executed on the device.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Including an instance of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html">VkVideoEncodeRateControlInfoKHR</a>
structure in the <code>pNext</code> chain of <code>pBeginInfo</code>
does not change the rate control state configured for the video session,
but only specifies the expected rate control state configured at the
time the recorded command is executed on the device which allows the
implementation to have information about the configured rate control
state at command buffer recording time. In order to change the current
rate control state of a video session, the application has to issue an
appropriate <a
href="vkCmdControlVideoCodingKHR.html">vkCmdControlVideoCodingKHR</a>
command as described in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-coding-control"
target="_blank" rel="noopener">Video Coding Control</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-state"
target="_blank" rel="noopener">Rate Control State</a> sections.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07231"
  id="VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07231"></a>
  VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07231  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support the video codec operation `pBeginInfo->videoSession` was
  created with, as returned by
  [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html)
  in
  [VkQueueFamilyVideoPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyVideoPropertiesKHR.html)::`videoCodecOperations`

- <a href="#VUID-vkCmdBeginVideoCodingKHR-None-07232"
  id="VUID-vkCmdBeginVideoCodingKHR-None-07232"></a>
  VUID-vkCmdBeginVideoCodingKHR-None-07232  
  There **must** be no <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a> queries

- <a href="#VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07233"
  id="VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07233"></a>
  VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07233  
  If `commandBuffer` is an unprotected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, then `pBeginInfo->videoSession` **must** not have been
  created with `VK_VIDEO_SESSION_CREATE_PROTECTED_CONTENT_BIT_KHR`

- <a href="#VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07234"
  id="VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07234"></a>
  VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07234  
  If `commandBuffer` is a protected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, then `pBeginInfo->videoSession` **must** have been
  created with `VK_VIDEO_SESSION_CREATE_PROTECTED_CONTENT_BIT_KHR`

- <a href="#VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07235"
  id="VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07235"></a>
  VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07235  
  If `commandBuffer` is an unprotected command buffer, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, and the `pPictureResource` member of any element of
  `pBeginInfo->pReferenceSlots` is not `NULL`, then
  `pPictureResource->imageViewBinding` for that element **must** not
  specify an image view created from a protected image

- <a href="#VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07236"
  id="VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07236"></a>
  VUID-vkCmdBeginVideoCodingKHR-commandBuffer-07236  
  If `commandBuffer` is a protected command buffer <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, and the `pPictureResource` member of any element of
  `pBeginInfo->pReferenceSlots` is not `NULL`, then
  `pPictureResource->imageViewBinding` for that element **must** specify
  an image view created from a protected image

- <a href="#VUID-vkCmdBeginVideoCodingKHR-slotIndex-07239"
  id="VUID-vkCmdBeginVideoCodingKHR-slotIndex-07239"></a>
  VUID-vkCmdBeginVideoCodingKHR-slotIndex-07239  
  If the `slotIndex` member of any element of
  `pBeginInfo->pReferenceSlots` is not negative, then it **must**
  specify the index of a DPB slot that is in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
  target="_blank" rel="noopener">active state</a> in
  `pBeginInfo->videoSession` at the time the command is executed on the
  device

- <a href="#VUID-vkCmdBeginVideoCodingKHR-pPictureResource-07265"
  id="VUID-vkCmdBeginVideoCodingKHR-pPictureResource-07265"></a>
  VUID-vkCmdBeginVideoCodingKHR-pPictureResource-07265  
  Each video picture resource specified by any non-`NULL`
  `pPictureResource` member specified in the elements of
  `pBeginInfo->pReferenceSlots` for which `slotIndex` is not negative
  **must** <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-matching"
  target="_blank" rel="noopener">match</a> one of the video picture
  resources currently associated with the DPB slot index of
  `pBeginInfo->videoSession` specified by `slotIndex` at the time the
  command is executed on the device

- <a href="#VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-08253"
  id="VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-08253"></a>
  VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-08253  
  If `pBeginInfo->videoSession` was created with a video encode
  operation and the `pNext` chain of `pBeginInfo` does not include an
  instance of the
  [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html)
  structure, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control mode</a> configured for
  `pBeginInfo->videoSession` at the time the command is executed on the
  device **must** be `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DEFAULT_KHR`

- <a href="#VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-08254"
  id="VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-08254"></a>
  VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-08254  
  If `pBeginInfo->videoSession` was created with a video encode
  operation and the `pNext` chain of `pBeginInfo` includes an instance
  of the
  [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html)
  structure, then it **must** <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-state-matching"
  target="_blank" rel="noopener">match</a> the rate control state
  configured for `pBeginInfo->videoSession` at the time the command is
  executed on the device

- <a href="#VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-08255"
  id="VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-08255"></a>
  VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-08255  
  If `pBeginInfo->videoSession` was created with the video codec
  operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, the current
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control mode</a> is not
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DEFAULT_KHR` or
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`, and
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`requiresGopRemainingFrames`
  is `VK_TRUE`, as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the `pBeginInfo->videoSession` was created with,
  then the `pNext` chain of `pBeginInfo` **must** include an instance of
  the
  [VkVideoEncodeH264GopRemainingFrameInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264GopRemainingFrameInfoKHR.html)
  with its `useGopRemainingFrames` member set to `VK_TRUE`

- <a href="#VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-08256"
  id="VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-08256"></a>
  VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-08256  
  If `pBeginInfo->videoSession` was created with the video codec
  operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, the current
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control mode</a> is not
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DEFAULT_KHR` or
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`, and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`requiresGopRemainingFrames`
  is `VK_TRUE`, as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile the `pBeginInfo->videoSession` was created with,
  then the `pNext` chain of `pBeginInfo` **must** include an instance of
  the
  [VkVideoEncodeH265GopRemainingFrameInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265GopRemainingFrameInfoKHR.html)
  with its `useGopRemainingFrames` member set to `VK_TRUE`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBeginVideoCodingKHR-commandBuffer-parameter"
  id="VUID-vkCmdBeginVideoCodingKHR-commandBuffer-parameter"></a>
  VUID-vkCmdBeginVideoCodingKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-parameter"
  id="VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-parameter"></a>
  VUID-vkCmdBeginVideoCodingKHR-pBeginInfo-parameter  
  `pBeginInfo` **must** be a valid pointer to a valid
  [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html) structure

- <a href="#VUID-vkCmdBeginVideoCodingKHR-commandBuffer-recording"
  id="VUID-vkCmdBeginVideoCodingKHR-commandBuffer-recording"></a>
  VUID-vkCmdBeginVideoCodingKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBeginVideoCodingKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdBeginVideoCodingKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdBeginVideoCodingKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support decode, or encode operations

- <a href="#VUID-vkCmdBeginVideoCodingKHR-renderpass"
  id="VUID-vkCmdBeginVideoCodingKHR-renderpass"></a>
  VUID-vkCmdBeginVideoCodingKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdBeginVideoCodingKHR-videocoding"
  id="VUID-vkCmdBeginVideoCodingKHR-videocoding"></a>
  VUID-vkCmdBeginVideoCodingKHR-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBeginVideoCodingKHR-bufferlevel"
  id="VUID-vkCmdBeginVideoCodingKHR-bufferlevel"></a>
  VUID-vkCmdBeginVideoCodingKHR-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr class="header">
<th class="tableblock halign-left valign-top"><a
href="#VkCommandBufferLevel">Command Buffer Levels</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginRenderPass">Render Pass Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginVideoCodingKHR">Video Coding Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#VkQueueFlagBits">Supported Queue Types</a></th>
<th class="tableblock halign-left valign-top"><a
href="#fundamentals-queueoperation-command-types">Command Type</a></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td class="tableblock halign-left valign-top"><p>Primary</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Decode<br />
Encode</p></td>
<td class="tableblock halign-left valign-top"><p>Action<br />
State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBeginVideoCodingKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
