# vkCmdSetVertexInputEXT(3) Manual Page

## Name

vkCmdSetVertexInputEXT - Set the vertex input state dynamically for a
command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the vertex input
attribute and vertex input binding descriptions, call:

``` c
// Provided by VK_EXT_shader_object, VK_EXT_vertex_input_dynamic_state
void vkCmdSetVertexInputEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    vertexBindingDescriptionCount,
    const VkVertexInputBindingDescription2EXT*  pVertexBindingDescriptions,
    uint32_t                                    vertexAttributeDescriptionCount,
    const VkVertexInputAttributeDescription2EXT* pVertexAttributeDescriptions);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `vertexBindingDescriptionCount` is the number of vertex binding
  descriptions provided in `pVertexBindingDescriptions`.

- `pVertexBindingDescriptions` is a pointer to an array of
  [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription2EXT.html)
  structures.

- `vertexAttributeDescriptionCount` is the number of vertex attribute
  descriptions provided in `pVertexAttributeDescriptions`.

- `pVertexAttributeDescriptions` is a pointer to an array of
  [VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription2EXT.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

This command sets the vertex input attribute and vertex input binding
descriptions state for subsequent drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`pVertexInputState`
values used to create the currently active pipeline.

If drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or if the bound
pipeline state object was also created with the
`VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE` dynamic state enabled,
then [vkCmdBindVertexBuffers2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindVertexBuffers2.html) can be used
instead of `vkCmdSetVertexInputEXT` to dynamically set the stride.

Valid Usage

- <a href="#VUID-vkCmdSetVertexInputEXT-None-08546"
  id="VUID-vkCmdSetVertexInputEXT-None-08546"></a>
  VUID-vkCmdSetVertexInputEXT-None-08546  
  Either the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-vertexInputDynamicState"
  target="_blank" rel="noopener"><code>vertexInputDynamicState</code></a>
  feature or the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderObject"
  target="_blank" rel="noopener"><code>shaderObject</code></a> feature
  or both **must** be enabled

- <a
  href="#VUID-vkCmdSetVertexInputEXT-vertexBindingDescriptionCount-04791"
  id="VUID-vkCmdSetVertexInputEXT-vertexBindingDescriptionCount-04791"></a>
  VUID-vkCmdSetVertexInputEXT-vertexBindingDescriptionCount-04791  
  `vertexBindingDescriptionCount` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxVertexInputBindings`

- <a
  href="#VUID-vkCmdSetVertexInputEXT-vertexAttributeDescriptionCount-04792"
  id="VUID-vkCmdSetVertexInputEXT-vertexAttributeDescriptionCount-04792"></a>
  VUID-vkCmdSetVertexInputEXT-vertexAttributeDescriptionCount-04792  
  `vertexAttributeDescriptionCount` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxVertexInputAttributes`

- <a href="#VUID-vkCmdSetVertexInputEXT-binding-04793"
  id="VUID-vkCmdSetVertexInputEXT-binding-04793"></a>
  VUID-vkCmdSetVertexInputEXT-binding-04793  
  For every `binding` specified by each element of
  `pVertexAttributeDescriptions`, a
  [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription2EXT.html)
  **must** exist in `pVertexBindingDescriptions` with the same value of
  `binding`

- <a href="#VUID-vkCmdSetVertexInputEXT-pVertexBindingDescriptions-04794"
  id="VUID-vkCmdSetVertexInputEXT-pVertexBindingDescriptions-04794"></a>
  VUID-vkCmdSetVertexInputEXT-pVertexBindingDescriptions-04794  
  All elements of `pVertexBindingDescriptions` **must** describe
  distinct binding numbers

- <a
  href="#VUID-vkCmdSetVertexInputEXT-pVertexAttributeDescriptions-04795"
  id="VUID-vkCmdSetVertexInputEXT-pVertexAttributeDescriptions-04795"></a>
  VUID-vkCmdSetVertexInputEXT-pVertexAttributeDescriptions-04795  
  All elements of `pVertexAttributeDescriptions` **must** describe
  distinct attribute locations

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetVertexInputEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetVertexInputEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetVertexInputEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetVertexInputEXT-pVertexBindingDescriptions-parameter"
  id="VUID-vkCmdSetVertexInputEXT-pVertexBindingDescriptions-parameter"></a>
  VUID-vkCmdSetVertexInputEXT-pVertexBindingDescriptions-parameter  
  If `vertexBindingDescriptionCount` is not `0`,
  `pVertexBindingDescriptions` **must** be a valid pointer to an array
  of `vertexBindingDescriptionCount` valid
  [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription2EXT.html)
  structures

- <a
  href="#VUID-vkCmdSetVertexInputEXT-pVertexAttributeDescriptions-parameter"
  id="VUID-vkCmdSetVertexInputEXT-pVertexAttributeDescriptions-parameter"></a>
  VUID-vkCmdSetVertexInputEXT-pVertexAttributeDescriptions-parameter  
  If `vertexAttributeDescriptionCount` is not `0`,
  `pVertexAttributeDescriptions` **must** be a valid pointer to an array
  of `vertexAttributeDescriptionCount` valid
  [VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription2EXT.html)
  structures

- <a href="#VUID-vkCmdSetVertexInputEXT-commandBuffer-recording"
  id="VUID-vkCmdSetVertexInputEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetVertexInputEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetVertexInputEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetVertexInputEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetVertexInputEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetVertexInputEXT-videocoding"
  id="VUID-vkCmdSetVertexInputEXT-videocoding"></a>
  VUID-vkCmdSetVertexInputEXT-videocoding  
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

[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VK_EXT_vertex_input_dynamic_state](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_vertex_input_dynamic_state.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription2EXT.html),
[VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription2EXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetVertexInputEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
