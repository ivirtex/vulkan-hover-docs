# vkCmdBindShadersEXT(3) Manual Page

## Name

vkCmdBindShadersEXT - Bind shader objects to a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Once shader objects have been created, they **can** be bound to the
command buffer using the command:

``` c
// Provided by VK_EXT_shader_object
void vkCmdBindShadersEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    stageCount,
    const VkShaderStageFlagBits*                pStages,
    const VkShaderEXT*                          pShaders);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer that the shader object will be
  bound to.

- `stageCount` is the length of the `pStages` and `pShaders` arrays.

- `pStages` is a pointer to an array of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) values specifying
  one stage per array index that is affected by the corresponding value
  in the `pShaders` array.

- `pShaders` is a pointer to an array of `VkShaderEXT` handles and/or
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) values describing the shader
  binding operations to be performed on each stage in `pStages`.

## <a href="#_description" class="anchor"></a>Description

When binding linked shaders, an application **may** bind them in any
combination of one or more calls to `vkCmdBindShadersEXT` (i.e., shaders
that were created linked together do not need to be bound in the same
`vkCmdBindShadersEXT` call).

Any shader object bound to a particular stage **may** be unbound by
setting its value in `pShaders` to
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html). If `pShaders` is `NULL`,
`vkCmdBindShadersEXT` behaves as if `pShaders` was an array of
`stageCount` [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) values (i.e., any
shaders bound to the stages specified in `pStages` are unbound).

Valid Usage

- <a href="#VUID-vkCmdBindShadersEXT-None-08462"
  id="VUID-vkCmdBindShadersEXT-None-08462"></a>
  VUID-vkCmdBindShadersEXT-None-08462  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderObject"
  target="_blank" rel="noopener"><code>shaderObject</code></a> feature
  **must** be enabled

- <a href="#VUID-vkCmdBindShadersEXT-pStages-08463"
  id="VUID-vkCmdBindShadersEXT-pStages-08463"></a>
  VUID-vkCmdBindShadersEXT-pStages-08463  
  Every element of `pStages` **must** be unique

- <a href="#VUID-vkCmdBindShadersEXT-pStages-08464"
  id="VUID-vkCmdBindShadersEXT-pStages-08464"></a>
  VUID-vkCmdBindShadersEXT-pStages-08464  
  `pStages` **must** not contain `VK_SHADER_STAGE_ALL_GRAPHICS` or
  `VK_SHADER_STAGE_ALL`

- <a href="#VUID-vkCmdBindShadersEXT-pStages-08465"
  id="VUID-vkCmdBindShadersEXT-pStages-08465"></a>
  VUID-vkCmdBindShadersEXT-pStages-08465  
  `pStages` **must** not contain `VK_SHADER_STAGE_RAYGEN_BIT_KHR`,
  `VK_SHADER_STAGE_ANY_HIT_BIT_KHR`,
  `VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR`, `VK_SHADER_STAGE_MISS_BIT_KHR`,
  `VK_SHADER_STAGE_INTERSECTION_BIT_KHR`, or
  `VK_SHADER_STAGE_CALLABLE_BIT_KHR`

- <a href="#VUID-vkCmdBindShadersEXT-pStages-08467"
  id="VUID-vkCmdBindShadersEXT-pStages-08467"></a>
  VUID-vkCmdBindShadersEXT-pStages-08467  
  `pStages` **must** not contain
  `VK_SHADER_STAGE_SUBPASS_SHADING_BIT_HUAWEI`

- <a href="#VUID-vkCmdBindShadersEXT-pStages-08468"
  id="VUID-vkCmdBindShadersEXT-pStages-08468"></a>
  VUID-vkCmdBindShadersEXT-pStages-08468  
  `pStages` **must** not contain
  `VK_SHADER_STAGE_CLUSTER_CULLING_BIT_HUAWEI`

- <a href="#VUID-vkCmdBindShadersEXT-pShaders-08469"
  id="VUID-vkCmdBindShadersEXT-pShaders-08469"></a>
  VUID-vkCmdBindShadersEXT-pShaders-08469  
  For each element of `pStages`, if `pShaders` is not `NULL`, and the
  element of the `pShaders` array with the same index is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it **must** have been created
  with a `stage` equal to the corresponding element of `pStages`

- <a href="#VUID-vkCmdBindShadersEXT-pShaders-08470"
  id="VUID-vkCmdBindShadersEXT-pShaders-08470"></a>
  VUID-vkCmdBindShadersEXT-pShaders-08470  
  If `pStages` contains both `VK_SHADER_STAGE_TASK_BIT_EXT` and
  `VK_SHADER_STAGE_VERTEX_BIT`, and `pShaders` is not `NULL`, and the
  same index in `pShaders` as `VK_SHADER_STAGE_TASK_BIT_EXT` in
  `pStages` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the same index
  in `pShaders` as `VK_SHADER_STAGE_VERTEX_BIT` in `pStages` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdBindShadersEXT-pShaders-08471"
  id="VUID-vkCmdBindShadersEXT-pShaders-08471"></a>
  VUID-vkCmdBindShadersEXT-pShaders-08471  
  If `pStages` contains both `VK_SHADER_STAGE_MESH_BIT_EXT` and
  `VK_SHADER_STAGE_VERTEX_BIT`, and `pShaders` is not `NULL`, and the
  same index in `pShaders` as `VK_SHADER_STAGE_MESH_BIT_EXT` in
  `pStages` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the same index
  in `pShaders` as `VK_SHADER_STAGE_VERTEX_BIT` in `pStages` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdBindShadersEXT-pShaders-08474"
  id="VUID-vkCmdBindShadersEXT-pShaders-08474"></a>
  VUID-vkCmdBindShadersEXT-pShaders-08474  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-tessellationShader"
  target="_blank" rel="noopener"><code>tessellationShader</code></a>
  feature is not enabled, and `pStages` contains
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` or
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, and `pShaders` is not
  `NULL`, the same index or indices in `pShaders` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdBindShadersEXT-pShaders-08475"
  id="VUID-vkCmdBindShadersEXT-pShaders-08475"></a>
  VUID-vkCmdBindShadersEXT-pShaders-08475  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-geometryShader"
  target="_blank" rel="noopener"><code>geometryShader</code></a> feature
  is not enabled, and `pStages` contains `VK_SHADER_STAGE_GEOMETRY_BIT`,
  and `pShaders` is not `NULL`, the same index in `pShaders` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdBindShadersEXT-pShaders-08490"
  id="VUID-vkCmdBindShadersEXT-pShaders-08490"></a>
  VUID-vkCmdBindShadersEXT-pShaders-08490  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-taskShader"
  target="_blank" rel="noopener"><code>taskShader</code></a> feature is
  not enabled, and `pStages` contains `VK_SHADER_STAGE_TASK_BIT_EXT`,
  and `pShaders` is not `NULL`, the same index in `pShaders` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdBindShadersEXT-pShaders-08491"
  id="VUID-vkCmdBindShadersEXT-pShaders-08491"></a>
  VUID-vkCmdBindShadersEXT-pShaders-08491  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-meshShader"
  target="_blank" rel="noopener"><code>meshShader</code></a> feature is
  not enabled, and `pStages` contains `VK_SHADER_STAGE_MESH_BIT_EXT`,
  and `pShaders` is not `NULL`, the same index in `pShaders` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdBindShadersEXT-pShaders-08476"
  id="VUID-vkCmdBindShadersEXT-pShaders-08476"></a>
  VUID-vkCmdBindShadersEXT-pShaders-08476  
  If `pStages` contains `VK_SHADER_STAGE_COMPUTE_BIT`, the
  `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdBindShadersEXT-pShaders-08477"
  id="VUID-vkCmdBindShadersEXT-pShaders-08477"></a>
  VUID-vkCmdBindShadersEXT-pShaders-08477  
  If `pStages` contains `VK_SHADER_STAGE_VERTEX_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`,
  `VK_SHADER_STAGE_GEOMETRY_BIT`, or `VK_SHADER_STAGE_FRAGMENT_BIT`, the
  `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBindShadersEXT-pShaders-08478"
  id="VUID-vkCmdBindShadersEXT-pShaders-08478"></a>
  VUID-vkCmdBindShadersEXT-pShaders-08478  
  If `pStages` contains `VK_SHADER_STAGE_MESH_BIT_EXT` or
  `VK_SHADER_STAGE_TASK_BIT_EXT`, the `VkCommandPool` that
  `commandBuffer` was allocated from **must** support graphics
  operations

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBindShadersEXT-commandBuffer-parameter"
  id="VUID-vkCmdBindShadersEXT-commandBuffer-parameter"></a>
  VUID-vkCmdBindShadersEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBindShadersEXT-pStages-parameter"
  id="VUID-vkCmdBindShadersEXT-pStages-parameter"></a>
  VUID-vkCmdBindShadersEXT-pStages-parameter  
  `pStages` **must** be a valid pointer to an array of `stageCount`
  valid [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) values

- <a href="#VUID-vkCmdBindShadersEXT-pShaders-parameter"
  id="VUID-vkCmdBindShadersEXT-pShaders-parameter"></a>
  VUID-vkCmdBindShadersEXT-pShaders-parameter  
  If `pShaders` is not `NULL`, `pShaders` **must** be a valid pointer to
  an array of `stageCount` valid or
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)
  handles

- <a href="#VUID-vkCmdBindShadersEXT-commandBuffer-recording"
  id="VUID-vkCmdBindShadersEXT-commandBuffer-recording"></a>
  VUID-vkCmdBindShadersEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBindShadersEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdBindShadersEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindShadersEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdBindShadersEXT-videocoding"
  id="VUID-vkCmdBindShadersEXT-videocoding"></a>
  VUID-vkCmdBindShadersEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBindShadersEXT-stageCount-arraylength"
  id="VUID-vkCmdBindShadersEXT-stageCount-arraylength"></a>
  VUID-vkCmdBindShadersEXT-stageCount-arraylength  
  `stageCount` **must** be greater than `0`

- <a href="#VUID-vkCmdBindShadersEXT-commonparent"
  id="VUID-vkCmdBindShadersEXT-commonparent"></a>
  VUID-vkCmdBindShadersEXT-commonparent  
  Both of `commandBuffer`, and the elements of `pShaders` that are valid
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
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html),
[VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindShadersEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
