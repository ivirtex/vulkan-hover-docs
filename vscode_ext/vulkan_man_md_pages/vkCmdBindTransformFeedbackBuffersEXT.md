# vkCmdBindTransformFeedbackBuffersEXT(3) Manual Page

## Name

vkCmdBindTransformFeedbackBuffersEXT - Bind transform feedback buffers
to a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To bind transform feedback buffers to a command buffer for use in
subsequent drawing commands, call:

``` c
// Provided by VK_EXT_transform_feedback
void vkCmdBindTransformFeedbackBuffersEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstBinding,
    uint32_t                                    bindingCount,
    const VkBuffer*                             pBuffers,
    const VkDeviceSize*                         pOffsets,
    const VkDeviceSize*                         pSizes);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `firstBinding` is the index of the first transform feedback binding
  whose state is updated by the command.

- `bindingCount` is the number of transform feedback bindings whose
  state is updated by the command.

- `pBuffers` is a pointer to an array of buffer handles.

- `pOffsets` is a pointer to an array of buffer offsets.

- `pSizes` is `NULL` or a pointer to an array of
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) buffer sizes, specifying the maximum
  number of bytes to capture to the corresponding transform feedback
  buffer. If `pSizes` is `NULL`, or the value of the `pSizes` array
  element is `VK_WHOLE_SIZE`, then the maximum number of bytes captured
  will be the size of the corresponding buffer minus the buffer offset.

## <a href="#_description" class="anchor"></a>Description

The values taken from elements i of `pBuffers`, `pOffsets` and `pSizes`
replace the current state for the transform feedback binding
`firstBinding` + i, for i in \[0, `bindingCount`). The transform
feedback binding is updated to start at the offset indicated by
`pOffsets`\[i\] from the start of the buffer `pBuffers`\[i\].

Valid Usage

- <a
  href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-transformFeedback-02355"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-transformFeedback-02355"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-transformFeedback-02355  
  `VkPhysicalDeviceTransformFeedbackFeaturesEXT`::`transformFeedback`
  **must** be enabled

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-firstBinding-02356"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-firstBinding-02356"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-firstBinding-02356  
  `firstBinding` **must** be less than
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBuffers`

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-firstBinding-02357"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-firstBinding-02357"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-firstBinding-02357  
  The sum of `firstBinding` and `bindingCount` **must** be less than or
  equal to
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBuffers`

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02358"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02358"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02358  
  All elements of `pOffsets` **must** be less than the size of the
  corresponding element in `pBuffers`

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02359"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02359"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02359  
  All elements of `pOffsets` **must** be a multiple of 4

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-02360"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-02360"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-02360  
  All elements of `pBuffers` **must** have been created with the
  `VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT` flag

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-pSize-02361"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-pSize-02361"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-pSize-02361  
  If the optional `pSize` array is specified, each element of `pSizes`
  **must** either be `VK_WHOLE_SIZE`, or be less than or equal to
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackBufferSize`

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-pSizes-02362"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-pSizes-02362"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-pSizes-02362  
  All elements of `pSizes` **must** be either `VK_WHOLE_SIZE`, or less
  than or equal to the size of the corresponding buffer in `pBuffers`

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02363"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02363"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-02363  
  All elements of `pOffsets` plus `pSizes`, where the `pSizes`, element
  is not `VK_WHOLE_SIZE`, **must** be less than or equal to the size of
  the corresponding buffer in `pBuffers`

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-02364"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-02364"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-02364  
  Each element of `pBuffers` that is non-sparse **must** be bound
  completely and contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-None-02365"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-None-02365"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-None-02365  
  Transform feedback **must** not be active when the
  `vkCmdBindTransformFeedbackBuffersEXT` command is recorded

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-parameter"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-parameter"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-parameter"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-parameter"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-pBuffers-parameter  
  `pBuffers` **must** be a valid pointer to an array of `bindingCount`
  valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handles

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-parameter"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-parameter"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-pOffsets-parameter  
  `pOffsets` **must** be a valid pointer to an array of `bindingCount`
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) values

- <a
  href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-recording"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-recording"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-videocoding"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-videocoding"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-bindingCount-arraylength"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-bindingCount-arraylength"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-bindingCount-arraylength  
  `bindingCount` **must** be greater than `0`

- <a href="#VUID-vkCmdBindTransformFeedbackBuffersEXT-commonparent"
  id="VUID-vkCmdBindTransformFeedbackBuffersEXT-commonparent"></a>
  VUID-vkCmdBindTransformFeedbackBuffersEXT-commonparent  
  Both of `commandBuffer`, and the elements of `pBuffers` **must** have
  been created, allocated, or retrieved from the same
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
<td class="tableblock halign-left valign-top"><p>Both</p></td>
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
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindTransformFeedbackBuffersEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
