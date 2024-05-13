# vkCmdSetDepthClampEnableEXT(3) Manual Page

## Name

vkCmdSetDepthClampEnableEXT - Specify dynamically whether depth clamping
is enabled in the command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> enable or disable
depth clamping, call:

``` c
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
void vkCmdSetDepthClampEnableEXT(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    depthClampEnable);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `depthClampEnable` specifies whether depth clamping is enabled.

## <a href="#_description" class="anchor"></a>Description

This command sets whether depth clamping is enabled or disabled for
subsequent drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_DEPTH_CLAMP_ENABLE_EXT` set
in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)::`depthClampEnable`
value used to create the currently active pipeline.

If the depth clamping state is changed dynamically, and the pipeline was
not created with `VK_DYNAMIC_STATE_DEPTH_CLIP_ENABLE_EXT` enabled, then
depth clipping is enabled when depth clamping is disabled and vice
versa.

Valid Usage

- <a href="#VUID-vkCmdSetDepthClampEnableEXT-None-09423"
  id="VUID-vkCmdSetDepthClampEnableEXT-None-09423"></a>
  VUID-vkCmdSetDepthClampEnableEXT-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3DepthClampEnable`](#features-extendedDynamicState3DepthClampEnable)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

- <a href="#VUID-vkCmdSetDepthClampEnableEXT-depthClamp-07449"
  id="VUID-vkCmdSetDepthClampEnableEXT-depthClamp-07449"></a>
  VUID-vkCmdSetDepthClampEnableEXT-depthClamp-07449  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-depthClamp"
  target="_blank" rel="noopener"><code>depthClamp</code></a> feature is
  not enabled, `depthClampEnable` must be `VK_FALSE`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-recording"
  id="VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetDepthClampEnableEXT-videocoding"
  id="VUID-vkCmdSetDepthClampEnableEXT-videocoding"></a>
  VUID-vkCmdSetDepthClampEnableEXT-videocoding  
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
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetDepthClampEnableEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
