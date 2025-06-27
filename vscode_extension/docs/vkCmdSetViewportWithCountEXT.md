# vkCmdSetViewportWithCount(3) Manual Page

## Name

vkCmdSetViewportWithCount - Set the viewport count and viewports
dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the viewport count
and viewports, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdSetViewportWithCount(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    viewportCount,
    const VkViewport*                           pViewports);
```

or the equivalent command

``` c
// Provided by VK_EXT_extended_dynamic_state, VK_EXT_shader_object
void vkCmdSetViewportWithCountEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    viewportCount,
    const VkViewport*                           pViewports);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `viewportCount` specifies the viewport count.

- `pViewports` specifies the viewports to use for drawing.

## <a href="#_description" class="anchor"></a>Description

This command sets the viewport count and viewports state for subsequent
drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the corresponding
[VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)::`viewportCount`
and `pViewports` values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetViewportWithCount-None-08971"
  id="VUID-vkCmdSetViewportWithCount-None-08971"></a>
  VUID-vkCmdSetViewportWithCount-None-08971  
  At least one of the following **must** be true:

  - the [`extendedDynamicState`](#features-extendedDynamicState) feature
    is enabled

  - the [`shaderObject`](#features-shaderObject) feature is enabled

  - the value of
    [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
    create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) parent of `commandBuffer`
    is greater than or equal to Version 1.3

- <a href="#VUID-vkCmdSetViewportWithCount-viewportCount-03394"
  id="VUID-vkCmdSetViewportWithCount-viewportCount-03394"></a>
  VUID-vkCmdSetViewportWithCount-viewportCount-03394  
  `viewportCount` **must** be between `1` and
  `VkPhysicalDeviceLimits`::`maxViewports`, inclusive

- <a href="#VUID-vkCmdSetViewportWithCount-viewportCount-03395"
  id="VUID-vkCmdSetViewportWithCount-viewportCount-03395"></a>
  VUID-vkCmdSetViewportWithCount-viewportCount-03395  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `viewportCount` **must** be `1`

- <a href="#VUID-vkCmdSetViewportWithCount-commandBuffer-04819"
  id="VUID-vkCmdSetViewportWithCount-commandBuffer-04819"></a>
  VUID-vkCmdSetViewportWithCount-commandBuffer-04819  
  `commandBuffer` **must** not have
  [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D`
  enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetViewportWithCount-commandBuffer-parameter"
  id="VUID-vkCmdSetViewportWithCount-commandBuffer-parameter"></a>
  VUID-vkCmdSetViewportWithCount-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetViewportWithCount-pViewports-parameter"
  id="VUID-vkCmdSetViewportWithCount-pViewports-parameter"></a>
  VUID-vkCmdSetViewportWithCount-pViewports-parameter  
  `pViewports` **must** be a valid pointer to an array of
  `viewportCount` valid [VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html) structures

- <a href="#VUID-vkCmdSetViewportWithCount-commandBuffer-recording"
  id="VUID-vkCmdSetViewportWithCount-commandBuffer-recording"></a>
  VUID-vkCmdSetViewportWithCount-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetViewportWithCount-commandBuffer-cmdpool"
  id="VUID-vkCmdSetViewportWithCount-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetViewportWithCount-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetViewportWithCount-videocoding"
  id="VUID-vkCmdSetViewportWithCount-videocoding"></a>
  VUID-vkCmdSetViewportWithCount-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdSetViewportWithCount-viewportCount-arraylength"
  id="VUID-vkCmdSetViewportWithCount-viewportCount-arraylength"></a>
  VUID-vkCmdSetViewportWithCount-viewportCount-arraylength  
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

[VK_EXT_extended_dynamic_state](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state.html),
[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetViewportWithCount"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
