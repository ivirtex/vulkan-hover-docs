# vkCmdSetViewportWScalingNV(3) Manual Page

## Name

vkCmdSetViewportWScalingNV - Set the viewport W scaling dynamically for
a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the viewport **W**
scaling parameters, call:

``` c
// Provided by VK_NV_clip_space_w_scaling
void vkCmdSetViewportWScalingNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstViewport,
    uint32_t                                    viewportCount,
    const VkViewportWScalingNV*                 pViewportWScalings);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `firstViewport` is the index of the first viewport whose parameters
  are updated by the command.

- `viewportCount` is the number of viewports whose parameters are
  updated by the command.

- `pViewportWScalings` is a pointer to an array of
  [VkViewportWScalingNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportWScalingNV.html) structures
  specifying viewport parameters.

## <a href="#_description" class="anchor"></a>Description

The viewport parameters taken from element i of `pViewportWScalings`
replace the current state for the viewport index `firstViewport` + i,
for i in \[0, `viewportCount`).

This command sets the viewport **W** scaling for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)::`pViewportWScalings`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetViewportWScalingNV-firstViewport-01324"
  id="VUID-vkCmdSetViewportWScalingNV-firstViewport-01324"></a>
  VUID-vkCmdSetViewportWScalingNV-firstViewport-01324  
  The sum of `firstViewport` and `viewportCount` **must** be between `1`
  and
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxViewports`,
  inclusive

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetViewportWScalingNV-commandBuffer-parameter"
  id="VUID-vkCmdSetViewportWScalingNV-commandBuffer-parameter"></a>
  VUID-vkCmdSetViewportWScalingNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetViewportWScalingNV-pViewportWScalings-parameter"
  id="VUID-vkCmdSetViewportWScalingNV-pViewportWScalings-parameter"></a>
  VUID-vkCmdSetViewportWScalingNV-pViewportWScalings-parameter  
  `pViewportWScalings` **must** be a valid pointer to an array of
  `viewportCount` [VkViewportWScalingNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportWScalingNV.html)
  structures

- <a href="#VUID-vkCmdSetViewportWScalingNV-commandBuffer-recording"
  id="VUID-vkCmdSetViewportWScalingNV-commandBuffer-recording"></a>
  VUID-vkCmdSetViewportWScalingNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetViewportWScalingNV-commandBuffer-cmdpool"
  id="VUID-vkCmdSetViewportWScalingNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetViewportWScalingNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetViewportWScalingNV-videocoding"
  id="VUID-vkCmdSetViewportWScalingNV-videocoding"></a>
  VUID-vkCmdSetViewportWScalingNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdSetViewportWScalingNV-viewportCount-arraylength"
  id="VUID-vkCmdSetViewportWScalingNV-viewportCount-arraylength"></a>
  VUID-vkCmdSetViewportWScalingNV-viewportCount-arraylength  
  `viewportCount` **must** be greater than `0`

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

[VK_NV_clip_space_w_scaling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_clip_space_w_scaling.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkViewportWScalingNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportWScalingNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetViewportWScalingNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
