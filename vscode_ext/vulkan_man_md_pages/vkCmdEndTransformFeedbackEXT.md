# vkCmdEndTransformFeedbackEXT(3) Manual Page

## Name

vkCmdEndTransformFeedbackEXT - Make transform feedback inactive in the
command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Transform feedback for specific transform feedback buffers is made
inactive by calling:

``` c
// Provided by VK_EXT_transform_feedback
void vkCmdEndTransformFeedbackEXT(
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
  [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handles to counter buffers. The counter
  buffers are used to record the current byte positions of each
  transform feedback buffer where the next vertex output data would be
  captured. This **can** be used by a subsequent
  [vkCmdBeginTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginTransformFeedbackEXT.html)
  call to resume transform feedback capture from this position. It can
  also be used by
  [vkCmdDrawIndirectByteCountEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirectByteCountEXT.html) to
  determine the vertex count of the draw call.

- `pCounterBufferOffsets` is `NULL` or a pointer to an array of
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) values specifying offsets within
  each of the `pCounterBuffers` where the counter values can be written.
  The location in each counter buffer at these offsets **must** be large
  enough to contain 4 bytes of data. The data stored at this location is
  the byte offset from the start of the transform feedback buffer
  binding where the next vertex data would be written. If
  `pCounterBufferOffsets` is `NULL`, then it is assumed the offsets are
  zero.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-transformFeedback-02374"
  id="VUID-vkCmdEndTransformFeedbackEXT-transformFeedback-02374"></a>
  VUID-vkCmdEndTransformFeedbackEXT-transformFeedback-02374  
  `VkPhysicalDeviceTransformFeedbackFeaturesEXT`::`transformFeedback`
  **must** be enabled

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-None-02375"
  id="VUID-vkCmdEndTransformFeedbackEXT-None-02375"></a>
  VUID-vkCmdEndTransformFeedbackEXT-None-02375  
  Transform feedback **must** be active

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-firstCounterBuffer-02376"
  id="VUID-vkCmdEndTransformFeedbackEXT-firstCounterBuffer-02376"></a>
  VUID-vkCmdEndTransformFeedbackEXT-firstCounterBuffer-02376  
  `firstCounterBuffer` **must** be less than
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBuffers`

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-firstCounterBuffer-02377"
  id="VUID-vkCmdEndTransformFeedbackEXT-firstCounterBuffer-02377"></a>
  VUID-vkCmdEndTransformFeedbackEXT-firstCounterBuffer-02377  
  The sum of `firstCounterBuffer` and `counterBufferCount` **must** be
  less than or equal to
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBuffers`

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-counterBufferCount-02608"
  id="VUID-vkCmdEndTransformFeedbackEXT-counterBufferCount-02608"></a>
  VUID-vkCmdEndTransformFeedbackEXT-counterBufferCount-02608  
  If `counterBufferCount` is not `0`, and `pCounterBuffers` is not
  `NULL`, `pCounterBuffers` **must** be a valid pointer to an array of
  `counterBufferCount` `VkBuffer` handles that are either valid or
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-pCounterBufferOffsets-02378"
  id="VUID-vkCmdEndTransformFeedbackEXT-pCounterBufferOffsets-02378"></a>
  VUID-vkCmdEndTransformFeedbackEXT-pCounterBufferOffsets-02378  
  For each buffer handle in the array, if it is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) it **must** reference a buffer
  large enough to hold 4 bytes at the corresponding offset from the
  `pCounterBufferOffsets` array

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-pCounterBuffer-02379"
  id="VUID-vkCmdEndTransformFeedbackEXT-pCounterBuffer-02379"></a>
  VUID-vkCmdEndTransformFeedbackEXT-pCounterBuffer-02379  
  If `pCounterBuffer` is `NULL`, then `pCounterBufferOffsets` **must**
  also be `NULL`

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-pCounterBuffers-02380"
  id="VUID-vkCmdEndTransformFeedbackEXT-pCounterBuffers-02380"></a>
  VUID-vkCmdEndTransformFeedbackEXT-pCounterBuffers-02380  
  For each buffer handle in the `pCounterBuffers` array that is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) it **must** have been created
  with a `usage` value containing
  `VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-parameter"
  id="VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-parameter"></a>
  VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdEndTransformFeedbackEXT-pCounterBufferOffsets-parameter"
  id="VUID-vkCmdEndTransformFeedbackEXT-pCounterBufferOffsets-parameter"></a>
  VUID-vkCmdEndTransformFeedbackEXT-pCounterBufferOffsets-parameter  
  If `counterBufferCount` is not `0`, and `pCounterBufferOffsets` is not
  `NULL`, `pCounterBufferOffsets` **must** be a valid pointer to an
  array of `counterBufferCount` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) values

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-recording"
  id="VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-recording"></a>
  VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdEndTransformFeedbackEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-renderpass"
  id="VUID-vkCmdEndTransformFeedbackEXT-renderpass"></a>
  VUID-vkCmdEndTransformFeedbackEXT-renderpass  
  This command **must** only be called inside of a render pass instance

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-videocoding"
  id="VUID-vkCmdEndTransformFeedbackEXT-videocoding"></a>
  VUID-vkCmdEndTransformFeedbackEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdEndTransformFeedbackEXT-commonparent"
  id="VUID-vkCmdEndTransformFeedbackEXT-commonparent"></a>
  VUID-vkCmdEndTransformFeedbackEXT-commonparent  
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
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdEndTransformFeedbackEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
