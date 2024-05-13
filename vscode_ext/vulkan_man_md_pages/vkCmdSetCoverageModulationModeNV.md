# vkCmdSetCoverageModulationModeNV(3) Manual Page

## Name

vkCmdSetCoverageModulationModeNV - Specify the coverage modulation mode
dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the
`coverageModulationMode` state, call:

``` c
// Provided by VK_EXT_extended_dynamic_state3 with VK_NV_framebuffer_mixed_samples, VK_EXT_shader_object with VK_NV_framebuffer_mixed_samples
void vkCmdSetCoverageModulationModeNV(
    VkCommandBuffer                             commandBuffer,
    VkCoverageModulationModeNV                  coverageModulationMode);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `coverageModulationMode` specifies the `coverageModulationMode` state.

## <a href="#_description" class="anchor"></a>Description

This command sets the `coverageModulationMode` state for subsequent
drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or the graphics
pipeline is created with `VK_DYNAMIC_STATE_COVERAGE_MODULATION_MODE_NV`
set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html)::`coverageModulationMode`
value used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetCoverageModulationModeNV-None-09423"
  id="VUID-vkCmdSetCoverageModulationModeNV-None-09423"></a>
  VUID-vkCmdSetCoverageModulationModeNV-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3CoverageModulationMode`](#features-extendedDynamicState3CoverageModulationMode)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-parameter"
  id="VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-parameter"></a>
  VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetCoverageModulationModeNV-coverageModulationMode-parameter"
  id="VUID-vkCmdSetCoverageModulationModeNV-coverageModulationMode-parameter"></a>
  VUID-vkCmdSetCoverageModulationModeNV-coverageModulationMode-parameter  
  `coverageModulationMode` **must** be a valid
  [VkCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageModulationModeNV.html) value

- <a href="#VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-recording"
  id="VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-recording"></a>
  VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-cmdpool"
  id="VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetCoverageModulationModeNV-videocoding"
  id="VUID-vkCmdSetCoverageModulationModeNV-videocoding"></a>
  VUID-vkCmdSetCoverageModulationModeNV-videocoding  
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

[VK_EXT_extended_dynamic_state3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state3.html),
[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VK_NV_framebuffer_mixed_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_framebuffer_mixed_samples.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageModulationModeNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetCoverageModulationModeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
