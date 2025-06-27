# vkCmdSetAlphaToOneEnableEXT(3) Manual Page

## Name

vkCmdSetAlphaToOneEnableEXT - Specify the alpha to one enable state
dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the
`alphaToOneEnable` state, call:

``` c
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
void vkCmdSetAlphaToOneEnableEXT(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    alphaToOneEnable);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `alphaToOneEnable` specifies the `alphaToOneEnable` state.

## <a href="#_description" class="anchor"></a>Description

This command sets the `alphaToOneEnable` state for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT` set
in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)::`alphaToOneEnable`
value used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetAlphaToOneEnableEXT-None-09423"
  id="VUID-vkCmdSetAlphaToOneEnableEXT-None-09423"></a>
  VUID-vkCmdSetAlphaToOneEnableEXT-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3AlphaToOneEnable`](#features-extendedDynamicState3AlphaToOneEnable)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

- <a href="#VUID-vkCmdSetAlphaToOneEnableEXT-alphaToOne-07607"
  id="VUID-vkCmdSetAlphaToOneEnableEXT-alphaToOne-07607"></a>
  VUID-vkCmdSetAlphaToOneEnableEXT-alphaToOne-07607  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-alphaToOne"
  target="_blank" rel="noopener"><code>alphaToOne</code></a> feature is
  not enabled, `alphaToOneEnable` **must** be `VK_FALSE`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-recording"
  id="VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetAlphaToOneEnableEXT-videocoding"
  id="VUID-vkCmdSetAlphaToOneEnableEXT-videocoding"></a>
  VUID-vkCmdSetAlphaToOneEnableEXT-videocoding  
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
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetAlphaToOneEnableEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
