# vkCmdSetDepthBoundsTestEnable(3) Manual Page

## Name

vkCmdSetDepthBoundsTestEnable - Set depth bounds test enable dynamically
for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically enable or disable</a> the
depth bounds test, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdSetDepthBoundsTestEnable(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    depthBoundsTestEnable);
```

or the equivalent command

``` c
// Provided by VK_EXT_extended_dynamic_state, VK_EXT_shader_object
void vkCmdSetDepthBoundsTestEnableEXT(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    depthBoundsTestEnable);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `depthBoundsTestEnable` specifies if the depth bounds test is enabled.

## <a href="#_description" class="anchor"></a>Description

This command sets the depth bounds enable for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE` set
in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)::`depthBoundsTestEnable`
value used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetDepthBoundsTestEnable-None-08971"
  id="VUID-vkCmdSetDepthBoundsTestEnable-None-08971"></a>
  VUID-vkCmdSetDepthBoundsTestEnable-None-08971  
  At least one of the following **must** be true:

  - the [`extendedDynamicState`](#features-extendedDynamicState) feature
    is enabled

  - the [`shaderObject`](#features-shaderObject) feature is enabled

  - the value of
    [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
    create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) parent of `commandBuffer`
    is greater than or equal to Version 1.3

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetDepthBoundsTestEnable-commandBuffer-parameter"
  id="VUID-vkCmdSetDepthBoundsTestEnable-commandBuffer-parameter"></a>
  VUID-vkCmdSetDepthBoundsTestEnable-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetDepthBoundsTestEnable-commandBuffer-recording"
  id="VUID-vkCmdSetDepthBoundsTestEnable-commandBuffer-recording"></a>
  VUID-vkCmdSetDepthBoundsTestEnable-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetDepthBoundsTestEnable-commandBuffer-cmdpool"
  id="VUID-vkCmdSetDepthBoundsTestEnable-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetDepthBoundsTestEnable-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetDepthBoundsTestEnable-videocoding"
  id="VUID-vkCmdSetDepthBoundsTestEnable-videocoding"></a>
  VUID-vkCmdSetDepthBoundsTestEnable-videocoding  
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
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetDepthBoundsTestEnable"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
