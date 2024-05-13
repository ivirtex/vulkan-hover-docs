# vkCmdBindVertexBuffers2(3) Manual Page

## Name

vkCmdBindVertexBuffers2 - Bind vertex buffers to a command buffer and
dynamically set strides



## <a href="#_c_specification" class="anchor"></a>C Specification

Alternatively, to bind vertex buffers, along with their sizes and
strides, to a command buffer for use in subsequent drawing commands,
call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdBindVertexBuffers2(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstBinding,
    uint32_t                                    bindingCount,
    const VkBuffer*                             pBuffers,
    const VkDeviceSize*                         pOffsets,
    const VkDeviceSize*                         pSizes,
    const VkDeviceSize*                         pStrides);
```

or the equivalent command

``` c
// Provided by VK_EXT_extended_dynamic_state, VK_EXT_shader_object
void vkCmdBindVertexBuffers2EXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstBinding,
    uint32_t                                    bindingCount,
    const VkBuffer*                             pBuffers,
    const VkDeviceSize*                         pOffsets,
    const VkDeviceSize*                         pSizes,
    const VkDeviceSize*                         pStrides);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `firstBinding` is the index of the first vertex input binding whose
  state is updated by the command.

- `bindingCount` is the number of vertex input bindings whose state is
  updated by the command.

- `pBuffers` is a pointer to an array of buffer handles.

- `pOffsets` is a pointer to an array of buffer offsets.

- `pSizes` is `NULL` or a pointer to an array of the size in bytes of
  vertex data bound from `pBuffers`.

- `pStrides` is `NULL` or a pointer to an array of buffer strides.

## <a href="#_description" class="anchor"></a>Description

The values taken from elements i of `pBuffers` and `pOffsets` replace
the current state for the vertex input binding `firstBinding` + i, for i
in \[0, `bindingCount`). The vertex input binding is updated to start at
the offset indicated by `pOffsets`\[i\] from the start of the buffer
`pBuffers`\[i\]. If `pSizes` is not `NULL` then `pSizes`\[i\] specifies
the bound size of the vertex buffer starting from the corresponding
elements of `pBuffers`\[i\] plus `pOffsets`\[i\]. If `pSizes`\[i\] is
`VK_WHOLE_SIZE` then the bound size is from `pBuffers`\[i\] plus
`pOffsets`\[i\] to the end of the buffer `pBuffers`\[i\]. All vertex
input attributes that use each of these bindings will use these updated
addresses in their address calculations for subsequent drawing commands.
If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
is enabled, elements of `pBuffers` **can** be
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and **can** be used by the vertex
shader. If a vertex input attribute is bound to a vertex input binding
that is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the values taken from
memory are considered to be zero, and missing G, B, or A components are
[filled with (0,0,1)](#fxvertex-input-extraction).

This command also <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically sets</a> the byte strides
between consecutive elements within buffer `pBuffers`\[i\] to the
corresponding `pStrides`\[i\] value when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE`
set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, strides are specified by the
[VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription.html)::`stride`
values used to create the currently active pipeline.

If drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a> or if the bound
pipeline state object was also created with the
`VK_DYNAMIC_STATE_VERTEX_INPUT_EXT` dynamic state enabled then
[vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetVertexInputEXT.html) **can** be used
instead of `vkCmdBindVertexBuffers2` to set the stride.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Unlike the static state to set the same, <code>pStrides</code> must
be between 0 and the maximum extent of the attributes in the binding. <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetVertexInputEXT.html">vkCmdSetVertexInputEXT</a> does not
have this restriction so can be used if other stride values are
desired.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkCmdBindVertexBuffers2-firstBinding-03355"
  id="VUID-vkCmdBindVertexBuffers2-firstBinding-03355"></a>
  VUID-vkCmdBindVertexBuffers2-firstBinding-03355  
  `firstBinding` **must** be less than
  `VkPhysicalDeviceLimits`::`maxVertexInputBindings`

- <a href="#VUID-vkCmdBindVertexBuffers2-firstBinding-03356"
  id="VUID-vkCmdBindVertexBuffers2-firstBinding-03356"></a>
  VUID-vkCmdBindVertexBuffers2-firstBinding-03356  
  The sum of `firstBinding` and `bindingCount` **must** be less than or
  equal to `VkPhysicalDeviceLimits`::`maxVertexInputBindings`

- <a href="#VUID-vkCmdBindVertexBuffers2-pOffsets-03357"
  id="VUID-vkCmdBindVertexBuffers2-pOffsets-03357"></a>
  VUID-vkCmdBindVertexBuffers2-pOffsets-03357  
  If `pSizes` is not `NULL`, all elements of `pOffsets` **must** be less
  than the size of the corresponding element in `pBuffers`

- <a href="#VUID-vkCmdBindVertexBuffers2-pSizes-03358"
  id="VUID-vkCmdBindVertexBuffers2-pSizes-03358"></a>
  VUID-vkCmdBindVertexBuffers2-pSizes-03358  
  If `pSizes` is not `NULL`, all elements of `pOffsets` plus `pSizes` ,
  where `pSizes` is not `VK_WHOLE_SIZE`, **must** be less than or equal
  to the size of the corresponding element in `pBuffers`

- <a href="#VUID-vkCmdBindVertexBuffers2-pBuffers-03359"
  id="VUID-vkCmdBindVertexBuffers2-pBuffers-03359"></a>
  VUID-vkCmdBindVertexBuffers2-pBuffers-03359  
  All elements of `pBuffers` **must** have been created with the
  `VK_BUFFER_USAGE_VERTEX_BUFFER_BIT` flag

- <a href="#VUID-vkCmdBindVertexBuffers2-pBuffers-03360"
  id="VUID-vkCmdBindVertexBuffers2-pBuffers-03360"></a>
  VUID-vkCmdBindVertexBuffers2-pBuffers-03360  
  Each element of `pBuffers` that is non-sparse **must** be bound
  completely and contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdBindVertexBuffers2-pBuffers-04111"
  id="VUID-vkCmdBindVertexBuffers2-pBuffers-04111"></a>
  VUID-vkCmdBindVertexBuffers2-pBuffers-04111  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
  target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
  is not enabled, all elements of `pBuffers` **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdBindVertexBuffers2-pBuffers-04112"
  id="VUID-vkCmdBindVertexBuffers2-pBuffers-04112"></a>
  VUID-vkCmdBindVertexBuffers2-pBuffers-04112  
  If an element of `pBuffers` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  then the corresponding element of `pOffsets` **must** be zero

- <a href="#VUID-vkCmdBindVertexBuffers2-pStrides-03362"
  id="VUID-vkCmdBindVertexBuffers2-pStrides-03362"></a>
  VUID-vkCmdBindVertexBuffers2-pStrides-03362  
  If `pStrides` is not `NULL` each element of `pStrides` **must** be
  less than or equal to
  `VkPhysicalDeviceLimits`::`maxVertexInputBindingStride`

- <a href="#VUID-vkCmdBindVertexBuffers2-pStrides-06209"
  id="VUID-vkCmdBindVertexBuffers2-pStrides-06209"></a>
  VUID-vkCmdBindVertexBuffers2-pStrides-06209  
  If `pStrides` is not `NULL` each element of `pStrides` **must** be
  either 0 or greater than or equal to the maximum extent of all vertex
  input attributes fetched from the corresponding binding, where the
  extent is calculated as the
  [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription.html)::`offset`
  plus
  [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription.html)::`format`
  size

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBindVertexBuffers2-commandBuffer-parameter"
  id="VUID-vkCmdBindVertexBuffers2-commandBuffer-parameter"></a>
  VUID-vkCmdBindVertexBuffers2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBindVertexBuffers2-pBuffers-parameter"
  id="VUID-vkCmdBindVertexBuffers2-pBuffers-parameter"></a>
  VUID-vkCmdBindVertexBuffers2-pBuffers-parameter  
  `pBuffers` **must** be a valid pointer to an array of `bindingCount`
  valid or [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handles

- <a href="#VUID-vkCmdBindVertexBuffers2-pOffsets-parameter"
  id="VUID-vkCmdBindVertexBuffers2-pOffsets-parameter"></a>
  VUID-vkCmdBindVertexBuffers2-pOffsets-parameter  
  `pOffsets` **must** be a valid pointer to an array of `bindingCount`
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) values

- <a href="#VUID-vkCmdBindVertexBuffers2-pSizes-parameter"
  id="VUID-vkCmdBindVertexBuffers2-pSizes-parameter"></a>
  VUID-vkCmdBindVertexBuffers2-pSizes-parameter  
  If `pSizes` is not `NULL`, `pSizes` **must** be a valid pointer to an
  array of `bindingCount` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) values

- <a href="#VUID-vkCmdBindVertexBuffers2-pStrides-parameter"
  id="VUID-vkCmdBindVertexBuffers2-pStrides-parameter"></a>
  VUID-vkCmdBindVertexBuffers2-pStrides-parameter  
  If `pStrides` is not `NULL`, `pStrides` **must** be a valid pointer to
  an array of `bindingCount` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) values

- <a href="#VUID-vkCmdBindVertexBuffers2-commandBuffer-recording"
  id="VUID-vkCmdBindVertexBuffers2-commandBuffer-recording"></a>
  VUID-vkCmdBindVertexBuffers2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBindVertexBuffers2-commandBuffer-cmdpool"
  id="VUID-vkCmdBindVertexBuffers2-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindVertexBuffers2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBindVertexBuffers2-videocoding"
  id="VUID-vkCmdBindVertexBuffers2-videocoding"></a>
  VUID-vkCmdBindVertexBuffers2-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBindVertexBuffers2-bindingCount-arraylength"
  id="VUID-vkCmdBindVertexBuffers2-bindingCount-arraylength"></a>
  VUID-vkCmdBindVertexBuffers2-bindingCount-arraylength  
  If any of `pSizes`, or `pStrides` are not `NULL`, `bindingCount`
  **must** be greater than `0`

- <a href="#VUID-vkCmdBindVertexBuffers2-commonparent"
  id="VUID-vkCmdBindVertexBuffers2-commonparent"></a>
  VUID-vkCmdBindVertexBuffers2-commonparent  
  Both of `commandBuffer`, and the elements of `pBuffers` that are valid
  handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindVertexBuffers2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
