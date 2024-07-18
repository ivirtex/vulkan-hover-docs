# vkCmdBeginTransformFeedbackEXT(3) Manual Page

## Name

vkCmdBeginTransformFeedbackEXT - Make transform feedback active in the
command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Transform feedback for specific transform feedback buffers is made
active by calling:

``` c
// Provided by VK_EXT_transform_feedback
void vkCmdBeginTransformFeedbackEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstCounterBuffer,
    uint32_t                                    counterBufferCount,
    const VkBuffer*                             pCounterBuffers,
    const VkDeviceSize*                         pCounterBufferOffsets);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `firstCounterBuffer` is the index of the first transform feedback
  buffer corresponding to `pCounterBuffers`\[0\] and
  `pCounterBufferOffsets`\[0\].

- `counterBufferCount` is the size of the `pCounterBuffers` and
  `pCounterBufferOffsets` arrays.

- `pCounterBuffers` is `NULL` or a pointer to an array of
  [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handles to counter buffers. Each buffer
  contains a 4 byte integer value representing the byte offset from the
  start of the corresponding transform feedback buffer from where to
  start capturing vertex data. If the byte offset stored to the counter
  buffer location was done using
  [vkCmdEndTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndTransformFeedbackEXT.html) it
  can be used to resume transform feedback from the previous location.
  If `pCounterBuffers` is `NULL`, then transform feedback will start
  capturing vertex data to byte offset zero in all bound transform
  feedback buffers. For each element of `pCounterBuffers` that is
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), transform feedback will start
  capturing vertex data to byte zero in the corresponding bound
  transform feedback buffer.

- `pCounterBufferOffsets` is `NULL` or a pointer to an array of
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) values specifying offsets within
  each of the `pCounterBuffers` where the counter values were previously
  written. The location in each counter buffer at these offsets **must**
  be large enough to contain 4 bytes of data. This data is the number of
  bytes captured by the previous transform feedback to this buffer. If
  `pCounterBufferOffsets` is `NULL`, then it is assumed the offsets are
  zero.

## <a href="#_description" class="anchor"></a>Description

The active transform feedback buffers will capture primitives emitted
from the corresponding `XfbBuffer` in the bound graphics pipeline. Any
`XfbBuffer` emitted that does not output to an active transform feedback
buffer will not be captured.

Valid Usage

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-transformFeedback-02366"
  id="VUID-vkCmdBeginTransformFeedbackEXT-transformFeedback-02366"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-transformFeedback-02366  
  `VkPhysicalDeviceTransformFeedbackFeaturesEXT`::`transformFeedback`
  **must** be enabled

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-None-02367"
  id="VUID-vkCmdBeginTransformFeedbackEXT-None-02367"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-None-02367  
  Transform feedback **must** not be active

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-02368"
  id="VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-02368"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-02368  
  `firstCounterBuffer` **must** be less than
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBuffers`

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-02369"
  id="VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-02369"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-02369  
  The sum of `firstCounterBuffer` and `counterBufferCount` **must** be
  less than or equal to
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBuffers`

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-counterBufferCount-02607"
  id="VUID-vkCmdBeginTransformFeedbackEXT-counterBufferCount-02607"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-counterBufferCount-02607  
  If `counterBufferCount` is not `0`, and `pCounterBuffers` is not
  `NULL`, `pCounterBuffers` **must** be a valid pointer to an array of
  `counterBufferCount` `VkBuffer` handles that are either valid or
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a
  href="#VUID-vkCmdBeginTransformFeedbackEXT-pCounterBufferOffsets-02370"
  id="VUID-vkCmdBeginTransformFeedbackEXT-pCounterBufferOffsets-02370"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-pCounterBufferOffsets-02370  
  For each buffer handle in the array, if it is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) it **must** reference a buffer
  large enough to hold 4 bytes at the corresponding offset from the
  `pCounterBufferOffsets` array

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-pCounterBuffer-02371"
  id="VUID-vkCmdBeginTransformFeedbackEXT-pCounterBuffer-02371"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-pCounterBuffer-02371  
  If `pCounterBuffer` is `NULL`, then `pCounterBufferOffsets` **must**
  also be `NULL`

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-pCounterBuffers-02372"
  id="VUID-vkCmdBeginTransformFeedbackEXT-pCounterBuffers-02372"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-pCounterBuffers-02372  
  For each buffer handle in the `pCounterBuffers` array that is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) it **must** have been created
  with a `usage` value containing
  `VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-09630"
  id="VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-09630"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-firstCounterBuffer-09630  
  The sum of `firstCounterBuffer` and `counterBufferCount` **must** be
  less than or equal to the number of transform feedback buffers
  currently bound by
  [vkCmdBindTransformFeedbackBuffersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindTransformFeedbackBuffersEXT.html)

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-None-06233"
  id="VUID-vkCmdBeginTransformFeedbackEXT-None-06233"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-None-06233  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderObject"
  target="_blank" rel="noopener"><code>shaderObject</code></a> feature
  is not enabled, a valid graphics pipeline **must** be bound to
  `VK_PIPELINE_BIND_POINT_GRAPHICS`

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-None-04128"
  id="VUID-vkCmdBeginTransformFeedbackEXT-None-04128"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-None-04128  
  The last <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader stage</a> of
  the bound graphics pipeline **must** have been declared with the `Xfb`
  execution mode

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-None-02373"
  id="VUID-vkCmdBeginTransformFeedbackEXT-None-02373"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-None-02373  
  Transform feedback **must** not be made active in a render pass
  instance with multiview enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-parameter"
  id="VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-parameter"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdBeginTransformFeedbackEXT-pCounterBufferOffsets-parameter"
  id="VUID-vkCmdBeginTransformFeedbackEXT-pCounterBufferOffsets-parameter"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-pCounterBufferOffsets-parameter  
  If `counterBufferCount` is not `0`, and `pCounterBufferOffsets` is not
  `NULL`, `pCounterBufferOffsets` **must** be a valid pointer to an
  array of `counterBufferCount` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) values

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-recording"
  id="VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-recording"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-renderpass"
  id="VUID-vkCmdBeginTransformFeedbackEXT-renderpass"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-renderpass  
  This command **must** only be called inside of a render pass instance

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-videocoding"
  id="VUID-vkCmdBeginTransformFeedbackEXT-videocoding"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBeginTransformFeedbackEXT-commonparent"
  id="VUID-vkCmdBeginTransformFeedbackEXT-commonparent"></a>
  VUID-vkCmdBeginTransformFeedbackEXT-commonparent  
  Both of `commandBuffer`, and the elements of `pCounterBuffers` that
  are valid handles of non-ignored parameters **must** have been
  created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
<tr>
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
<tr>
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Inside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_transform_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_transform_feedback.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBeginTransformFeedbackEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
