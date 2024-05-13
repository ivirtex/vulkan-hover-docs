# vkCmdSetFrontFace(3) Manual Page

## Name

vkCmdSetFrontFace - Set front face orientation dynamically for a command
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the front face
orientation, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdSetFrontFace(
    VkCommandBuffer                             commandBuffer,
    VkFrontFace                                 frontFace);
```

or the equivalent command

``` c
// Provided by VK_EXT_extended_dynamic_state, VK_EXT_shader_object
void vkCmdSetFrontFaceEXT(
    VkCommandBuffer                             commandBuffer,
    VkFrontFace                                 frontFace);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `frontFace` is a [VkFrontFace](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrontFace.html) value specifying the
  front-facing triangle orientation to be used for culling.

## <a href="#_description" class="anchor"></a>Description

This command sets the front face orientation for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_FRONT_FACE` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)::`frontFace`
value used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetFrontFace-None-08971"
  id="VUID-vkCmdSetFrontFace-None-08971"></a>
  VUID-vkCmdSetFrontFace-None-08971  
  At least one of the following **must** be true:

  - the [`extendedDynamicState`](#features-extendedDynamicState) feature
    is enabled

  - the [`shaderObject`](#features-shaderObject) feature is enabled

  - the value of
    [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
    create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) parent of `commandBuffer`
    is greater than or equal to Version 1.3

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetFrontFace-commandBuffer-parameter"
  id="VUID-vkCmdSetFrontFace-commandBuffer-parameter"></a>
  VUID-vkCmdSetFrontFace-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetFrontFace-frontFace-parameter"
  id="VUID-vkCmdSetFrontFace-frontFace-parameter"></a>
  VUID-vkCmdSetFrontFace-frontFace-parameter  
  `frontFace` **must** be a valid [VkFrontFace](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrontFace.html) value

- <a href="#VUID-vkCmdSetFrontFace-commandBuffer-recording"
  id="VUID-vkCmdSetFrontFace-commandBuffer-recording"></a>
  VUID-vkCmdSetFrontFace-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetFrontFace-commandBuffer-cmdpool"
  id="VUID-vkCmdSetFrontFace-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetFrontFace-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetFrontFace-videocoding"
  id="VUID-vkCmdSetFrontFace-videocoding"></a>
  VUID-vkCmdSetFrontFace-videocoding  
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

[VK_EXT_extended_dynamic_state](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state.html),
[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkFrontFace](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrontFace.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetFrontFace"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
