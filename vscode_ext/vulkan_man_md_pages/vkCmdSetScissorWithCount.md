# vkCmdSetScissorWithCount(3) Manual Page

## Name

vkCmdSetScissorWithCount - Set the scissor count and scissor rectangular
bounds dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the scissor count and
scissor rectangular bounds, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdSetScissorWithCount(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    scissorCount,
    const VkRect2D*                             pScissors);
```

or the equivalent command

``` c
// Provided by VK_EXT_extended_dynamic_state, VK_EXT_shader_object
void vkCmdSetScissorWithCountEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    scissorCount,
    const VkRect2D*                             pScissors);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `scissorCount` specifies the scissor count.

- `pScissors` specifies the scissors to use for drawing.

## <a href="#_description" class="anchor"></a>Description

This command sets the scissor count and scissor rectangular bounds state
for subsequent drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the corresponding
[VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)::`scissorCount`
and `pScissors` values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetScissorWithCount-None-08971"
  id="VUID-vkCmdSetScissorWithCount-None-08971"></a>
  VUID-vkCmdSetScissorWithCount-None-08971  
  At least one of the following **must** be true:

  - the [`extendedDynamicState`](#features-extendedDynamicState) feature
    is enabled

  - the [`shaderObject`](#features-shaderObject) feature is enabled

  - the value of
    [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
    create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) parent of `commandBuffer`
    is greater than or equal to Version 1.3

- <a href="#VUID-vkCmdSetScissorWithCount-scissorCount-03397"
  id="VUID-vkCmdSetScissorWithCount-scissorCount-03397"></a>
  VUID-vkCmdSetScissorWithCount-scissorCount-03397  
  `scissorCount` **must** be between `1` and
  `VkPhysicalDeviceLimits`::`maxViewports`, inclusive

- <a href="#VUID-vkCmdSetScissorWithCount-scissorCount-03398"
  id="VUID-vkCmdSetScissorWithCount-scissorCount-03398"></a>
  VUID-vkCmdSetScissorWithCount-scissorCount-03398  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `scissorCount` **must** be `1`

- <a href="#VUID-vkCmdSetScissorWithCount-x-03399"
  id="VUID-vkCmdSetScissorWithCount-x-03399"></a>
  VUID-vkCmdSetScissorWithCount-x-03399  
  The `x` and `y` members of `offset` member of any element of
  `pScissors` **must** be greater than or equal to `0`

- <a href="#VUID-vkCmdSetScissorWithCount-offset-03400"
  id="VUID-vkCmdSetScissorWithCount-offset-03400"></a>
  VUID-vkCmdSetScissorWithCount-offset-03400  
  Evaluation of (`offset.x` + `extent.width`) **must** not cause a
  signed integer addition overflow for any element of `pScissors`

- <a href="#VUID-vkCmdSetScissorWithCount-offset-03401"
  id="VUID-vkCmdSetScissorWithCount-offset-03401"></a>
  VUID-vkCmdSetScissorWithCount-offset-03401  
  Evaluation of (`offset.y` + `extent.height`) **must** not cause a
  signed integer addition overflow for any element of `pScissors`

- <a href="#VUID-vkCmdSetScissorWithCount-commandBuffer-04820"
  id="VUID-vkCmdSetScissorWithCount-commandBuffer-04820"></a>
  VUID-vkCmdSetScissorWithCount-commandBuffer-04820  
  `commandBuffer` **must** not have
  [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D`
  enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetScissorWithCount-commandBuffer-parameter"
  id="VUID-vkCmdSetScissorWithCount-commandBuffer-parameter"></a>
  VUID-vkCmdSetScissorWithCount-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetScissorWithCount-pScissors-parameter"
  id="VUID-vkCmdSetScissorWithCount-pScissors-parameter"></a>
  VUID-vkCmdSetScissorWithCount-pScissors-parameter  
  `pScissors` **must** be a valid pointer to an array of `scissorCount`
  [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures

- <a href="#VUID-vkCmdSetScissorWithCount-commandBuffer-recording"
  id="VUID-vkCmdSetScissorWithCount-commandBuffer-recording"></a>
  VUID-vkCmdSetScissorWithCount-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetScissorWithCount-commandBuffer-cmdpool"
  id="VUID-vkCmdSetScissorWithCount-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetScissorWithCount-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetScissorWithCount-videocoding"
  id="VUID-vkCmdSetScissorWithCount-videocoding"></a>
  VUID-vkCmdSetScissorWithCount-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdSetScissorWithCount-scissorCount-arraylength"
  id="VUID-vkCmdSetScissorWithCount-scissorCount-arraylength"></a>
  VUID-vkCmdSetScissorWithCount-scissorCount-arraylength  
  `scissorCount` **must** be greater than `0`

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
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetScissorWithCount"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
