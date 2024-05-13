# vkCreateShadersEXT(3) Manual Page

## Name

vkCreateShadersEXT - Create one or more new shaders



## <a href="#_c_specification" class="anchor"></a>C Specification

To create one or more shader objects, call:

``` c
// Provided by VK_EXT_shader_object
VkResult vkCreateShadersEXT(
    VkDevice                                    device,
    uint32_t                                    createInfoCount,
    const VkShaderCreateInfoEXT*                pCreateInfos,
    const VkAllocationCallbacks*                pAllocator,
    VkShaderEXT*                                pShaders);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the shader objects.

- `createInfoCount` is the length of the `pCreateInfos` and `pShaders`
  arrays.

- `pCreateInfos` is a pointer to an array of
  [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html) structures.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pShaders` is a pointer to an array of [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)
  handles in which the resulting shader objects are returned.

## <a href="#_description" class="anchor"></a>Description

When this function returns, whether or not it succeeds, it is guaranteed
that every element of `pShaders` will have been overwritten by either
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a valid `VkShaderEXT` handle.

This means that whenever shader creation fails, the application **can**
determine which shader the returned error pertains to by locating the
first [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) element in `pShaders`. It
also means that an application **can** reliably clean up from a failed
call by iterating over the `pShaders` array and destroying every element
that is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html).

Valid Usage

- <a href="#VUID-vkCreateShadersEXT-None-08400"
  id="VUID-vkCreateShadersEXT-None-08400"></a>
  VUID-vkCreateShadersEXT-None-08400  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderObject"
  target="_blank" rel="noopener"><code>shaderObject</code></a> feature
  **must** be enabled

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-08402"
  id="VUID-vkCreateShadersEXT-pCreateInfos-08402"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-08402  
  If the `flags` member of any element of `pCreateInfos` includes
  `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`, the `flags` member of all other
  elements of `pCreateInfos` whose `stage` is
  `VK_SHADER_STAGE_VERTEX_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`,
  `VK_SHADER_STAGE_GEOMETRY_BIT`, or `VK_SHADER_STAGE_FRAGMENT_BIT`
  **must** also include `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-08403"
  id="VUID-vkCreateShadersEXT-pCreateInfos-08403"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-08403  
  If the `flags` member of any element of `pCreateInfos` includes
  `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`, the `flags` member of all other
  elements of `pCreateInfos` whose `stage` is
  `VK_SHADER_STAGE_TASK_BIT_EXT` or `VK_SHADER_STAGE_MESH_BIT_EXT`
  **must** also include `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-08404"
  id="VUID-vkCreateShadersEXT-pCreateInfos-08404"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-08404  
  If the `flags` member of any element of `pCreateInfos` whose `stage`
  is `VK_SHADER_STAGE_TASK_BIT_EXT` or `VK_SHADER_STAGE_MESH_BIT_EXT`
  includes `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`, there **must** be no
  member of `pCreateInfos` whose `stage` is `VK_SHADER_STAGE_VERTEX_BIT`
  and whose `flags` member includes
  `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-08405"
  id="VUID-vkCreateShadersEXT-pCreateInfos-08405"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-08405  
  If there is any element of `pCreateInfos` whose `stage` is
  `VK_SHADER_STAGE_MESH_BIT_EXT` and whose `flags` member includes both
  `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT` and
  `VK_SHADER_CREATE_NO_TASK_SHADER_BIT_EXT`, there **must** be no
  element of `pCreateInfos` whose `stage` is
  `VK_SHADER_STAGE_TASK_BIT_EXT` and whose `flags` member includes
  `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-08409"
  id="VUID-vkCreateShadersEXT-pCreateInfos-08409"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-08409  
  For each element of `pCreateInfos` whose `flags` member includes
  `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`, if there is any other element
  of `pCreateInfos` whose `stage` is logically later than the `stage` of
  the former and whose `flags` member also includes
  `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`, the `nextStage` of the former
  **must** be equal to the `stage` of the element with the logically
  earliest `stage` following the `stage` of the former whose `flags`
  member also includes `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-08410"
  id="VUID-vkCreateShadersEXT-pCreateInfos-08410"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-08410  
  The `stage` member of each element of `pCreateInfos` whose `flags`
  member includes `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT` **must** be
  unique

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-08411"
  id="VUID-vkCreateShadersEXT-pCreateInfos-08411"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-08411  
  The `codeType` member of all elements of `pCreateInfos` whose `flags`
  member includes `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT` **must** be the
  same

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-08867"
  id="VUID-vkCreateShadersEXT-pCreateInfos-08867"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-08867  
  If `pCreateInfos` contains elements with both
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` and
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, both elements' `flags`
  include `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`, both elements'
  `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` stage’s `pCode` contains an
  `OpExecutionMode` instruction specifying the type of subdivision, it
  **must** match the subdivision type specified in the
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT` stage

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-08868"
  id="VUID-vkCreateShadersEXT-pCreateInfos-08868"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-08868  
  If `pCreateInfos` contains elements with both
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` and
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, both elements' `flags`
  include `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`, both elements'
  `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` stage’s `pCode` contains an
  `OpExecutionMode` instruction specifying the orientation of triangles,
  it **must** match the triangle orientation specified in the
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT` stage

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-08869"
  id="VUID-vkCreateShadersEXT-pCreateInfos-08869"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-08869  
  If `pCreateInfos` contains elements with both
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` and
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, both elements' `flags`
  include `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`, both elements'
  `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` stage’s `pCode` contains an
  `OpExecutionMode` instruction specifying `PointMode`, the
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT` stage **must** also
  contain an `OpExecutionMode` instruction specifying `PointMode`

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-08870"
  id="VUID-vkCreateShadersEXT-pCreateInfos-08870"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-08870  
  If `pCreateInfos` contains elements with both
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` and
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, both elements' `flags`
  include `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`, both elements'
  `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` stage’s `pCode` contains an
  `OpExecutionMode` instruction specifying the spacing of segments on
  the edges of tessellated primitives, it **must** match the segment
  spacing specified in the `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`
  stage

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-08871"
  id="VUID-vkCreateShadersEXT-pCreateInfos-08871"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-08871  
  If `pCreateInfos` contains elements with both
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` and
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, both elements' `flags`
  include `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`, both elements'
  `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` stage’s `pCode` contains an
  `OpExecutionMode` instruction specifying the output patch size, it
  **must** match the output patch size specified in the
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT` stage

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-09632"
  id="VUID-vkCreateShadersEXT-pCreateInfos-09632"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-09632  
  If `pCreateInfos` contains a `VK_SHADER_STAGE_MESH_BIT_EXT` with
  `codeType` of `VK_SHADER_CODE_TYPE_SPIRV_EXT` and
  `VK_SHADER_CREATE_NO_TASK_SHADER_BIT_EXT` is not set, then the mesh
  shader’s entry point **must** not declare a variable with a
  `DrawIndex` `BuiltIn` decoration

Valid Usage (Implicit)

- <a href="#VUID-vkCreateShadersEXT-device-parameter"
  id="VUID-vkCreateShadersEXT-device-parameter"></a>
  VUID-vkCreateShadersEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateShadersEXT-pCreateInfos-parameter"
  id="VUID-vkCreateShadersEXT-pCreateInfos-parameter"></a>
  VUID-vkCreateShadersEXT-pCreateInfos-parameter  
  `pCreateInfos` **must** be a valid pointer to an array of
  `createInfoCount` valid
  [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html) structures

- <a href="#VUID-vkCreateShadersEXT-pAllocator-parameter"
  id="VUID-vkCreateShadersEXT-pAllocator-parameter"></a>
  VUID-vkCreateShadersEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateShadersEXT-pShaders-parameter"
  id="VUID-vkCreateShadersEXT-pShaders-parameter"></a>
  VUID-vkCreateShadersEXT-pShaders-parameter  
  `pShaders` **must** be a valid pointer to an array of
  `createInfoCount` [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) handles

- <a href="#VUID-vkCreateShadersEXT-createInfoCount-arraylength"
  id="VUID-vkCreateShadersEXT-createInfoCount-arraylength"></a>
  VUID-vkCreateShadersEXT-createInfoCount-arraylength  
  `createInfoCount` **must** be greater than `0`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPATIBLE_SHADER_BINARY_EXT`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html),
[VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateShadersEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
