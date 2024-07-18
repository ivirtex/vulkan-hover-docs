# VkPhysicalDeviceMeshShaderPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceMeshShaderPropertiesEXT - Structure describing mesh
shading properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMeshShaderPropertiesEXT` structure is defined as:

``` c
// Provided by VK_EXT_mesh_shader
typedef struct VkPhysicalDeviceMeshShaderPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxTaskWorkGroupTotalCount;
    uint32_t           maxTaskWorkGroupCount[3];
    uint32_t           maxTaskWorkGroupInvocations;
    uint32_t           maxTaskWorkGroupSize[3];
    uint32_t           maxTaskPayloadSize;
    uint32_t           maxTaskSharedMemorySize;
    uint32_t           maxTaskPayloadAndSharedMemorySize;
    uint32_t           maxMeshWorkGroupTotalCount;
    uint32_t           maxMeshWorkGroupCount[3];
    uint32_t           maxMeshWorkGroupInvocations;
    uint32_t           maxMeshWorkGroupSize[3];
    uint32_t           maxMeshSharedMemorySize;
    uint32_t           maxMeshPayloadAndSharedMemorySize;
    uint32_t           maxMeshOutputMemorySize;
    uint32_t           maxMeshPayloadAndOutputMemorySize;
    uint32_t           maxMeshOutputComponents;
    uint32_t           maxMeshOutputVertices;
    uint32_t           maxMeshOutputPrimitives;
    uint32_t           maxMeshOutputLayers;
    uint32_t           maxMeshMultiviewViewCount;
    uint32_t           meshOutputPerVertexGranularity;
    uint32_t           meshOutputPerPrimitiveGranularity;
    uint32_t           maxPreferredTaskWorkGroupInvocations;
    uint32_t           maxPreferredMeshWorkGroupInvocations;
    VkBool32           prefersLocalInvocationVertexOutput;
    VkBool32           prefersLocalInvocationPrimitiveOutput;
    VkBool32           prefersCompactVertexOutput;
    VkBool32           prefersCompactPrimitiveOutput;
} VkPhysicalDeviceMeshShaderPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

The members of the `VkPhysicalDeviceMeshShaderPropertiesEXT` structure
describe the following implementation-dependent limits:

## <a href="#_description" class="anchor"></a>Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-maxTaskWorkGroupTotalCount"></span>
  `maxTaskWorkGroupTotalCount` is the maximum number of total local
  workgroups that **can** be launched by a single mesh tasks drawing
  command. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-mesh-shading"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-mesh-shading</a>.

- <span id="limits-maxTaskWorkGroupCount"></span>
  `maxTaskWorkGroupCount`\[3\] is the maximum number of local workgroups
  that **can** be launched by a single mesh tasks drawing command. These
  three values represent the maximum number of local workgroups for the
  X, Y, and Z dimensions, respectively. The workgroup count parameters
  to the drawing commands **must** be less than or equal to the
  corresponding limit. The product of these dimensions **must** be less
  than or equal to `maxTaskWorkGroupTotalCount`.

- <span id="limits-maxTaskWorkGroupInvocations"></span>
  `maxTaskWorkGroupInvocations` is the maximum total number of task
  shader invocations in a single local workgroup. The product of the X,
  Y, and Z sizes, as specified by the `LocalSize` or `LocalSizeId`
  execution mode in shader modules or by the object decorated by the
  `WorkgroupSize` decoration, **must** be less than or equal to this
  limit.

- <span id="limits-maxTaskWorkGroupSize"></span>
  `maxTaskWorkGroupSize`\[3\] is the maximum size of a local task
  workgroup, per dimension. These three values represent the maximum
  local workgroup size in the X, Y, and Z dimensions, respectively. The
  `x`, `y`, and `z` sizes, as specified by the `LocalSize` or
  `LocalSizeId` execution mode or by the object decorated by the
  `WorkgroupSize` decoration in shader modules, **must** be less than or
  equal to the corresponding limit.

- <span id="limits-maxTaskPayloadSize"></span> `maxTaskPayloadSize` is
  the maximum total storage size, in bytes, available for variables
  declared with the `TaskPayloadWorkgroupEXT` storage class in shader
  modules in the task shader stage.

- <span id="limits-maxTaskSharedMemorySize"></span>
  `maxTaskSharedMemorySize` is the maximum total storage size, in bytes,
  available for variables declared with the `Workgroup` storage class in
  shader modules in the task shader stage.

- <span id="limits-maxTaskPayloadAndSharedMemorySize"></span>
  `maxTaskPayloadAndSharedMemorySize` is the maximum total storage size,
  in bytes, available for variables that are declared with the
  `TaskPayloadWorkgroupEXT` or `Workgroup` storage class, in shader
  modules in the task shader stage.

- <span id="limits-maxMeshWorkGroupTotalCount"></span>
  `maxMeshWorkGroupTotalCount` is the maximum number of local output
  tasks a single task shader workgroup can emit.

- <span id="limits-maxMeshWorkGroupCount"></span>
  `maxMeshWorkGroupCount`\[3\] is the maximum number of local output
  tasks a single task shader workgroup can emit, per dimension. These
  three values represent the maximum number of local output tasks for
  the X, Y, and Z dimensions, respectively. The workgroup count
  parameters to the `OpEmitMeshTasksEXT` **must** be less than or equal
  to the corresponding limit. The product of these dimensions **must**
  be less than or equal to `maxMeshWorkGroupTotalCount`.

- <span id="limits-maxMeshWorkGroupInvocations"></span>
  `maxMeshWorkGroupInvocations` is the maximum total number of mesh
  shader invocations in a single local workgroup. The product of the X,
  Y, and Z sizes, as specified by the `LocalSize` or `LocalSizeId`
  execution mode in shader modules or by the object decorated by the
  `WorkgroupSize` decoration, **must** be less than or equal to this
  limit.

- <span id="limits-maxMeshWorkGroupSize"></span>
  `maxMeshWorkGroupSize`\[3\] is the maximum size of a local mesh
  workgroup, per dimension. These three values represent the maximum
  local workgroup size in the X, Y, and Z dimensions, respectively. The
  `x`, `y`, and `z` sizes, as specified by the `LocalSize` or
  `LocalSizeId` execution mode or by the object decorated by the
  `WorkgroupSize` decoration in shader modules, **must** be less than or
  equal to the corresponding limit.

- <span id="limits-maxMeshSharedMemorySize"></span>
  `maxMeshSharedMemorySize` is the maximum total storage size, in bytes,
  available for variables declared with the `Workgroup` storage class in
  shader modules in the mesh shader stage.

- <span id="limits-maxMeshPayloadAndSharedMemorySize"></span>
  `maxMeshPayloadAndSharedMemorySize` is the maximum total storage size,
  in bytes, available for variables that are declared with the
  `TaskPayloadWorkgroupEXT` or `Workgroup` storage class in shader
  modules in the mesh shader stage.

- <span id="limits-maxMeshOutputMemorySize"></span>
  `maxMeshOutputMemorySize` is the maximum total storage size, in bytes,
  available for output variables in shader modules in the mesh shader
  stage, according to the formula in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#mesh-output"
  target="_blank" rel="noopener">Mesh Shader Output</a>.

- <span id="limits-maxMeshPayloadAndOutputMemorySize"></span>
  `maxMeshPayloadAndOutputMemorySize` is the maximum total storage size,
  in bytes, available for variables that are declared with the
  `TaskPayloadWorkgroupEXT` storage class, or output variables in shader
  modules in the mesh shader stage, according to the formula in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#mesh-output"
  target="_blank" rel="noopener">Mesh Shader Output</a>.

- <span id="limits-maxMeshOutputComponents"></span>
  `maxMeshOutputComponents` is the maximum number of components of
  output variables which **can** be output from the mesh shader stage.

- <span id="limits-maxMeshOutputVertices"></span>
  `maxMeshOutputVertices` is the maximum number of vertices which
  **can** be emitted by a single mesh shader workgroup.

- <span id="limits-maxMeshOutputPrimitives"></span>
  `maxMeshOutputPrimitives` is the maximum number of primitives which
  **can** be emitted by a single mesh shader workgroup.

- <span id="limits-maxMeshOutputLayers"></span> `maxMeshOutputLayers` is
  one greater than the maximum layer index that **can** be output from
  the mesh shader stage.

- <span id="limits-maxMeshMultiviewViewCount"></span>
  `maxMeshMultiviewViewCount` is one greater than the maximum view index
  that **can** be used by any mesh shader.

- <span id="limits-meshOutputPerVertexGranularity"></span>
  `meshOutputPerVertexGranularity` is the granularity of vertex
  allocation. The number of output vertices allocated for the mesh
  shader stage is padded to a multiple of this number. The value can be
  used to calculate the required storage size for output variables in
  shader modules in the mesh shader stage, which **must** be less than
  or equal to `maxMeshOutputMemorySize`.

- <span id="limits-meshOutputPerPrimitiveGranularity"></span>
  `meshOutputPerPrimitiveGranularity` is the granularity of primitive
  allocation. The number of output primitives allocated for the mesh
  shader stage is padded to a multiple of this number. The value can be
  used to calculate the required storage size for output variables in
  shader modules in the mesh shader stage, which **must** be less than
  or equal to `maxMeshOutputMemorySize`.

- <span id="limits-maxPreferredTaskWorkGroupInvocations"></span>
  `maxPreferredTaskWorkGroupInvocations` is the maximum number of task
  shader invocations in a single workgroup that is preferred by the
  implementation for optimal performance. The value is guaranteed to be
  a multiple of a supported subgroup size for the task shader stage.

- <span id="limits-maxPreferredMeshWorkGroupInvocations"></span>
  `maxPreferredMeshWorkGroupInvocations` is the maximum number of mesh
  shader invocations in a single workgroup that is preferred by the
  implementation for optimal performance. The value is guaranteed to be
  a multiple of a supported subgroup size for the mesh shader stage.

- <span id="limits-prefersLocalInvocationVertexOutput"></span>
  `prefersLocalInvocationVertexOutput` specifies whether writes to the
  vertex output array in a mesh shader yield best performance when the
  array index matches `LocalInvocationIndex`.

- <span id="limits-prefersLocalInvocationPrimitiveOutput"></span>
  `prefersLocalInvocationPrimitiveOutput` specifies whether writes to
  the primitive output array in a mesh shader yield best performance
  when the array index matches `LocalInvocationIndex`.

- <span id="limits-prefersCompactVertexOutput"></span>
  `prefersCompactVertexOutput` specifies whether output vertices should
  be compacted after custom culling in the mesh shader for best
  performance, otherwise keeping the vertices at their original location
  may be better.

- <span id="limits-prefersCompactPrimitiveOutput"></span>
  `prefersCompactPrimitiveOutput` specifies whether output primitives
  should be compacted after custom culling in the mesh shader for best
  performance, otherwise the use of `CullPrimitiveEXT` may be better.

If the `VkPhysicalDeviceMeshShaderPropertiesEXT` structure is included
in the `pNext` chain of
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html), it is
filled with the implementation-dependent limits.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMeshShaderPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceMeshShaderPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceMeshShaderPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_mesh_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_mesh_shader.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMeshShaderPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
