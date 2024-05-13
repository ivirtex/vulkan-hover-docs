# VkPhysicalDeviceMeshShaderPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceMeshShaderPropertiesNV - Structure describing mesh
shading properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMeshShaderPropertiesNV` structure is defined as:

``` c
// Provided by VK_NV_mesh_shader
typedef struct VkPhysicalDeviceMeshShaderPropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxDrawMeshTasksCount;
    uint32_t           maxTaskWorkGroupInvocations;
    uint32_t           maxTaskWorkGroupSize[3];
    uint32_t           maxTaskTotalMemorySize;
    uint32_t           maxTaskOutputCount;
    uint32_t           maxMeshWorkGroupInvocations;
    uint32_t           maxMeshWorkGroupSize[3];
    uint32_t           maxMeshTotalMemorySize;
    uint32_t           maxMeshOutputVertices;
    uint32_t           maxMeshOutputPrimitives;
    uint32_t           maxMeshMultiviewViewCount;
    uint32_t           meshOutputPerVertexGranularity;
    uint32_t           meshOutputPerPrimitiveGranularity;
} VkPhysicalDeviceMeshShaderPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxDrawMeshTasksCount` is the maximum number of local workgroups that
  **can** be launched by a single draw mesh tasks command. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-mesh-shading"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-mesh-shading</a>.

- `maxTaskWorkGroupInvocations` is the maximum total number of task
  shader invocations in a single local workgroup. The product of the X,
  Y, and Z sizes, as specified by the `LocalSize` or `LocalSizeId`
  execution mode in shader modules or by the object decorated by the
  `WorkgroupSize` decoration, **must** be less than or equal to this
  limit.

- `maxTaskWorkGroupSize`\[3\] is the maximum size of a local task
  workgroup. These three values represent the maximum local workgroup
  size in the X, Y, and Z dimensions, respectively. The `x`, `y`, and
  `z` sizes, as specified by the `LocalSize` or `LocalSizeId` execution
  mode or by the object decorated by the `WorkgroupSize` decoration in
  shader modules, **must** be less than or equal to the corresponding
  limit.

- `maxTaskTotalMemorySize` is the maximum number of bytes that the task
  shader can use in total for shared and output memory combined.

- `maxTaskOutputCount` is the maximum number of output tasks a single
  task shader workgroup can emit.

- `maxMeshWorkGroupInvocations` is the maximum total number of mesh
  shader invocations in a single local workgroup. The product of the X,
  Y, and Z sizes, as specified by the `LocalSize` or `LocalSizeId`
  execution mode in shader modules or by the object decorated by the
  `WorkgroupSize` decoration, **must** be less than or equal to this
  limit.

- `maxMeshWorkGroupSize`\[3\] is the maximum size of a local mesh
  workgroup. These three values represent the maximum local workgroup
  size in the X, Y, and Z dimensions, respectively. The `x`, `y`, and
  `z` sizes, as specified by the `LocalSize` or `LocalSizeId` execution
  mode or by the object decorated by the `WorkgroupSize` decoration in
  shader modules, **must** be less than or equal to the corresponding
  limit.

- `maxMeshTotalMemorySize` is the maximum number of bytes that the mesh
  shader can use in total for shared and output memory combined.

- `maxMeshOutputVertices` is the maximum number of vertices a mesh
  shader output can store.

- `maxMeshOutputPrimitives` is the maximum number of primitives a mesh
  shader output can store.

- `maxMeshMultiviewViewCount` is the maximum number of multiview views a
  mesh shader can use.

- `meshOutputPerVertexGranularity` is the granularity with which mesh
  vertex outputs are allocated. The value can be used to compute the
  memory size used by the mesh shader, which must be less than or equal
  to `maxMeshTotalMemorySize`.

- `meshOutputPerPrimitiveGranularity` is the granularity with which mesh
  outputs qualified as per-primitive are allocated. The value can be
  used to compute the memory size used by the mesh shader, which must be
  less than or equal to `maxMeshTotalMemorySize`.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMeshShaderPropertiesNV` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMeshShaderPropertiesNV-sType-sType"
  id="VUID-VkPhysicalDeviceMeshShaderPropertiesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceMeshShaderPropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_PROPERTIES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_mesh_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_mesh_shader.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMeshShaderPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
