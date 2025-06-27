# vkCmdSetViewportSwizzleNV(3) Manual Page

## Name

vkCmdSetViewportSwizzleNV - Specify the viewport swizzle state
dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the viewport swizzle
state, call:

``` c
// Provided by VK_EXT_extended_dynamic_state3 with VK_NV_viewport_swizzle, VK_EXT_shader_object with VK_NV_viewport_swizzle
void vkCmdSetViewportSwizzleNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstViewport,
    uint32_t                                    viewportCount,
    const VkViewportSwizzleNV*                  pViewportSwizzles);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `firstViewport` is the index of the first viewport whose parameters
  are updated by the command.

- `viewportCount` is the number of viewports whose parameters are
  updated by the command.

- `pViewportSwizzles` is a pointer to an array of
  [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportSwizzleNV.html) structures specifying
  viewport swizzles.

## <a href="#_description" class="anchor"></a>Description

This command sets the viewport swizzle state for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_VIEWPORT_SWIZZLE_NV` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html)::`viewportCount`,
and
[VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html)::`pViewportSwizzles`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetViewportSwizzleNV-None-09423"
  id="VUID-vkCmdSetViewportSwizzleNV-None-09423"></a>
  VUID-vkCmdSetViewportSwizzleNV-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3ViewportSwizzle`](#features-extendedDynamicState3ViewportSwizzle)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetViewportSwizzleNV-commandBuffer-parameter"
  id="VUID-vkCmdSetViewportSwizzleNV-commandBuffer-parameter"></a>
  VUID-vkCmdSetViewportSwizzleNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetViewportSwizzleNV-pViewportSwizzles-parameter"
  id="VUID-vkCmdSetViewportSwizzleNV-pViewportSwizzles-parameter"></a>
  VUID-vkCmdSetViewportSwizzleNV-pViewportSwizzles-parameter  
  `pViewportSwizzles` **must** be a valid pointer to an array of
  `viewportCount` valid [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportSwizzleNV.html)
  structures

- <a href="#VUID-vkCmdSetViewportSwizzleNV-commandBuffer-recording"
  id="VUID-vkCmdSetViewportSwizzleNV-commandBuffer-recording"></a>
  VUID-vkCmdSetViewportSwizzleNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetViewportSwizzleNV-commandBuffer-cmdpool"
  id="VUID-vkCmdSetViewportSwizzleNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetViewportSwizzleNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetViewportSwizzleNV-videocoding"
  id="VUID-vkCmdSetViewportSwizzleNV-videocoding"></a>
  VUID-vkCmdSetViewportSwizzleNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdSetViewportSwizzleNV-viewportCount-arraylength"
  id="VUID-vkCmdSetViewportSwizzleNV-viewportCount-arraylength"></a>
  VUID-vkCmdSetViewportSwizzleNV-viewportCount-arraylength  
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
[VK_NV_viewport_swizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_viewport_swizzle.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportSwizzleNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetViewportSwizzleNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
