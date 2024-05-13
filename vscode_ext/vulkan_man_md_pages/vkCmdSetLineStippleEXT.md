# vkCmdSetLineStippleKHR(3) Manual Page

## Name

vkCmdSetLineStippleKHR - Set line stipple dynamically for a command
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the line stipple
state, call:

``` c
// Provided by VK_KHR_line_rasterization
void vkCmdSetLineStippleKHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    lineStippleFactor,
    uint16_t                                    lineStipplePattern);
```

or the equivalent command

``` c
// Provided by VK_EXT_line_rasterization
void vkCmdSetLineStippleEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    lineStippleFactor,
    uint16_t                                    lineStipplePattern);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `lineStippleFactor` is the repeat factor used in stippled line
  rasterization.

- `lineStipplePattern` is the bit pattern used in stippled line
  rasterization.

## <a href="#_description" class="anchor"></a>Description

This command sets the line stipple state for subsequent drawing commands
when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_LINE_STIPPLE_EXT` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)::`lineStippleFactor`
and
[VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)::`lineStipplePattern`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetLineStippleKHR-lineStippleFactor-02776"
  id="VUID-vkCmdSetLineStippleKHR-lineStippleFactor-02776"></a>
  VUID-vkCmdSetLineStippleKHR-lineStippleFactor-02776  
  `lineStippleFactor` **must** be in the range \[1,256\]

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetLineStippleKHR-commandBuffer-parameter"
  id="VUID-vkCmdSetLineStippleKHR-commandBuffer-parameter"></a>
  VUID-vkCmdSetLineStippleKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetLineStippleKHR-commandBuffer-recording"
  id="VUID-vkCmdSetLineStippleKHR-commandBuffer-recording"></a>
  VUID-vkCmdSetLineStippleKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetLineStippleKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdSetLineStippleKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetLineStippleKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetLineStippleKHR-videocoding"
  id="VUID-vkCmdSetLineStippleKHR-videocoding"></a>
  VUID-vkCmdSetLineStippleKHR-videocoding  
  This command **must** only be called outside of a video coding scope

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

[VK_EXT_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_line_rasterization.html),
[VK_KHR_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_line_rasterization.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetLineStippleKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
