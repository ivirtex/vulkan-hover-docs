# vkCmdSetStencilOp(3) Manual Page

## Name

vkCmdSetStencilOp - Set stencil operation dynamically for a command
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the stencil
operation, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdSetStencilOp(
    VkCommandBuffer                             commandBuffer,
    VkStencilFaceFlags                          faceMask,
    VkStencilOp                                 failOp,
    VkStencilOp                                 passOp,
    VkStencilOp                                 depthFailOp,
    VkCompareOp                                 compareOp);
```

or the equivalent command

``` c
// Provided by VK_EXT_extended_dynamic_state, VK_EXT_shader_object
void vkCmdSetStencilOpEXT(
    VkCommandBuffer                             commandBuffer,
    VkStencilFaceFlags                          faceMask,
    VkStencilOp                                 failOp,
    VkStencilOp                                 passOp,
    VkStencilOp                                 depthFailOp,
    VkCompareOp                                 compareOp);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `faceMask` is a bitmask of
  [VkStencilFaceFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilFaceFlagBits.html) specifying the set
  of stencil state for which to update the stencil operation.

- `failOp` is a [VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html) value specifying the
  action performed on samples that fail the stencil test.

- `passOp` is a [VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html) value specifying the
  action performed on samples that pass both the depth and stencil
  tests.

- `depthFailOp` is a [VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html) value specifying
  the action performed on samples that pass the stencil test and fail
  the depth test.

- `compareOp` is a [VkCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html) value specifying the
  comparison operator used in the stencil test.

## <a href="#_description" class="anchor"></a>Description

This command sets the stencil operation for subsequent drawing commands
when when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_STENCIL_OP` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the corresponding
`VkPipelineDepthStencilStateCreateInfo`::`failOp`, `passOp`,
`depthFailOp`, and `compareOp` values used to create the currently
active pipeline, for both front and back faces.

Valid Usage

- <a href="#VUID-vkCmdSetStencilOp-None-08971"
  id="VUID-vkCmdSetStencilOp-None-08971"></a>
  VUID-vkCmdSetStencilOp-None-08971  
  At least one of the following **must** be true:

  - the [`extendedDynamicState`](#features-extendedDynamicState) feature
    is enabled

  - the [`shaderObject`](#features-shaderObject) feature is enabled

  - the value of
    [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
    create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) parent of `commandBuffer`
    is greater than or equal to Version 1.3

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetStencilOp-commandBuffer-parameter"
  id="VUID-vkCmdSetStencilOp-commandBuffer-parameter"></a>
  VUID-vkCmdSetStencilOp-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetStencilOp-faceMask-parameter"
  id="VUID-vkCmdSetStencilOp-faceMask-parameter"></a>
  VUID-vkCmdSetStencilOp-faceMask-parameter  
  `faceMask` **must** be a valid combination of
  [VkStencilFaceFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilFaceFlagBits.html) values

- <a href="#VUID-vkCmdSetStencilOp-faceMask-requiredbitmask"
  id="VUID-vkCmdSetStencilOp-faceMask-requiredbitmask"></a>
  VUID-vkCmdSetStencilOp-faceMask-requiredbitmask  
  `faceMask` **must** not be `0`

- <a href="#VUID-vkCmdSetStencilOp-failOp-parameter"
  id="VUID-vkCmdSetStencilOp-failOp-parameter"></a>
  VUID-vkCmdSetStencilOp-failOp-parameter  
  `failOp` **must** be a valid [VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html) value

- <a href="#VUID-vkCmdSetStencilOp-passOp-parameter"
  id="VUID-vkCmdSetStencilOp-passOp-parameter"></a>
  VUID-vkCmdSetStencilOp-passOp-parameter  
  `passOp` **must** be a valid [VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html) value

- <a href="#VUID-vkCmdSetStencilOp-depthFailOp-parameter"
  id="VUID-vkCmdSetStencilOp-depthFailOp-parameter"></a>
  VUID-vkCmdSetStencilOp-depthFailOp-parameter  
  `depthFailOp` **must** be a valid [VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html)
  value

- <a href="#VUID-vkCmdSetStencilOp-compareOp-parameter"
  id="VUID-vkCmdSetStencilOp-compareOp-parameter"></a>
  VUID-vkCmdSetStencilOp-compareOp-parameter  
  `compareOp` **must** be a valid [VkCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html) value

- <a href="#VUID-vkCmdSetStencilOp-commandBuffer-recording"
  id="VUID-vkCmdSetStencilOp-commandBuffer-recording"></a>
  VUID-vkCmdSetStencilOp-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetStencilOp-commandBuffer-cmdpool"
  id="VUID-vkCmdSetStencilOp-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetStencilOp-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetStencilOp-videocoding"
  id="VUID-vkCmdSetStencilOp-videocoding"></a>
  VUID-vkCmdSetStencilOp-videocoding  
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

[VK_EXT_extended_dynamic_state](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state.html),
[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html),
[VkStencilFaceFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilFaceFlags.html),
[VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetStencilOp"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
