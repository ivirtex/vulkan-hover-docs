# vkCmdSetRepresentativeFragmentTestEnableNV(3) Manual Page

## Name

vkCmdSetRepresentativeFragmentTestEnableNV - Specify the representative
fragment test enable dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the
`representativeFragmentTestEnable` state, call:

``` c
// Provided by VK_EXT_extended_dynamic_state3 with VK_NV_representative_fragment_test, VK_EXT_shader_object with VK_NV_representative_fragment_test
void vkCmdSetRepresentativeFragmentTestEnableNV(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    representativeFragmentTestEnable);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `representativeFragmentTestEnable` specifies the
  `representativeFragmentTestEnable` state.

## <a href="#_description" class="anchor"></a>Description

This command sets the `representativeFragmentTestEnable` state for
subsequent drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with
`VK_DYNAMIC_STATE_REPRESENTATIVE_FRAGMENT_TEST_ENABLE_NV` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineRepresentativeFragmentTestStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRepresentativeFragmentTestStateCreateInfoNV.html)::`representativeFragmentTestEnable`
value used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetRepresentativeFragmentTestEnableNV-None-09423"
  id="VUID-vkCmdSetRepresentativeFragmentTestEnableNV-None-09423"></a>
  VUID-vkCmdSetRepresentativeFragmentTestEnableNV-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3RepresentativeFragmentTestEnable`](#features-extendedDynamicState3RepresentativeFragmentTestEnable)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-parameter"
  id="VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-parameter"></a>
  VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-recording"
  id="VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-recording"></a>
  VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-cmdpool"
  id="VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetRepresentativeFragmentTestEnableNV-videocoding"
  id="VUID-vkCmdSetRepresentativeFragmentTestEnableNV-videocoding"></a>
  VUID-vkCmdSetRepresentativeFragmentTestEnableNV-videocoding  
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

[VK_EXT_extended_dynamic_state3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state3.html),
[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VK_NV_representative_fragment_test](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_representative_fragment_test.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetRepresentativeFragmentTestEnableNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
