# vkCmdSetBlendConstants(3) Manual Page

## Name

vkCmdSetBlendConstants - Set the values of blend constants



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set and change</a> the blend
constants, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdSetBlendConstants(
    VkCommandBuffer                             commandBuffer,
    const float                                 blendConstants[4]);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `blendConstants` is a pointer to an array of four values specifying
  the R<sub>c</sub>, G<sub>c</sub>, B<sub>c</sub>, and A<sub>c</sub>
  components of the blend constant color used in blending, depending on
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blendfactors"
  target="_blank" rel="noopener">blend factor</a>.

## <a href="#_description" class="anchor"></a>Description

This command sets blend constants for subsequent drawing commands when
when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or the graphics
pipeline is created with `VK_DYNAMIC_STATE_BLEND_CONSTANTS` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)::`blendConstants`
values used to create the currently active pipeline.

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetBlendConstants-commandBuffer-parameter"
  id="VUID-vkCmdSetBlendConstants-commandBuffer-parameter"></a>
  VUID-vkCmdSetBlendConstants-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetBlendConstants-commandBuffer-recording"
  id="VUID-vkCmdSetBlendConstants-commandBuffer-recording"></a>
  VUID-vkCmdSetBlendConstants-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetBlendConstants-commandBuffer-cmdpool"
  id="VUID-vkCmdSetBlendConstants-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetBlendConstants-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetBlendConstants-videocoding"
  id="VUID-vkCmdSetBlendConstants-videocoding"></a>
  VUID-vkCmdSetBlendConstants-videocoding  
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

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetBlendConstants"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
