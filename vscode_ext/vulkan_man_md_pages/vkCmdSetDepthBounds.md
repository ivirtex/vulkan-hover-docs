# vkCmdSetDepthBounds(3) Manual Page

## Name

vkCmdSetDepthBounds - Set depth bounds range dynamically for a command
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the depth bounds
range, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdSetDepthBounds(
    VkCommandBuffer                             commandBuffer,
    float                                       minDepthBounds,
    float                                       maxDepthBounds);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `minDepthBounds` is the minimum depth bound.

- `maxDepthBounds` is the maximum depth bound.

## <a href="#_description" class="anchor"></a>Description

This command sets the depth bounds range for subsequent drawing commands
when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_DEPTH_BOUNDS` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)::`minDepthBounds`
and
[VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)::`maxDepthBounds`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetDepthBounds-minDepthBounds-00600"
  id="VUID-vkCmdSetDepthBounds-minDepthBounds-00600"></a>
  VUID-vkCmdSetDepthBounds-minDepthBounds-00600  
  If the
  [`VK_EXT_depth_range_unrestricted`](VK_EXT_depth_range_unrestricted.html)
  extension is not enabled `minDepthBounds` **must** be between `0.0`
  and `1.0`, inclusive

- <a href="#VUID-vkCmdSetDepthBounds-maxDepthBounds-00601"
  id="VUID-vkCmdSetDepthBounds-maxDepthBounds-00601"></a>
  VUID-vkCmdSetDepthBounds-maxDepthBounds-00601  
  If the
  [`VK_EXT_depth_range_unrestricted`](VK_EXT_depth_range_unrestricted.html)
  extension is not enabled `maxDepthBounds` **must** be between `0.0`
  and `1.0`, inclusive

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetDepthBounds-commandBuffer-parameter"
  id="VUID-vkCmdSetDepthBounds-commandBuffer-parameter"></a>
  VUID-vkCmdSetDepthBounds-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetDepthBounds-commandBuffer-recording"
  id="VUID-vkCmdSetDepthBounds-commandBuffer-recording"></a>
  VUID-vkCmdSetDepthBounds-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetDepthBounds-commandBuffer-cmdpool"
  id="VUID-vkCmdSetDepthBounds-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetDepthBounds-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetDepthBounds-videocoding"
  id="VUID-vkCmdSetDepthBounds-videocoding"></a>
  VUID-vkCmdSetDepthBounds-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetDepthBounds"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
