# VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI(3) Manual Page

## Name

VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI - Structure
describing cluster culling shader properties supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI` structure is
defined as:

``` c
// Provided by VK_HUAWEI_cluster_culling_shader
typedef struct VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxWorkGroupCount[3];
    uint32_t           maxWorkGroupSize[3];
    uint32_t           maxOutputClusterCount;
    VkDeviceSize       indirectBufferOffsetAlignment;
} VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxWorkGroupCount`\[3\] is the maximum number of local workgroups
  that can be launched by a single command. These three value represent
  the maximum local workgroup count in the X, Y and Z dimensions,
  respectively. In the current implementation, the values of Y and Z are
  both implicitly set as one. groupCountX of DrawCluster command must be
  less than or equal to maxWorkGroupCount\[0\].

- `maxWorkGroupSize`\[3\] is the maximum size of a local workgroup.
  These three value represent the maximum local workgroup size in the X,
  Y and Z dimensions, respectively. The x, y and z sizes, as specified
  by the `LocalSize` or `LocalSizeId` execution mode or by the object
  decorated by the WorkgroupSize decoration in shader modules, must be
  less than or equal to the corresponding limit. In the current
  implementation, the maximum workgroup size of the X dimension is 32,
  the others are 1.

- `maxOutputClusterCount` is the maximum number of output cluster a
  single cluster culling shader workgroup can emit.

- `indirectBufferOffsetAlignment` indicates the alignment for cluster
  drawing command buffer stride.
  [vkCmdDrawClusterIndirectHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawClusterIndirectHUAWEI.html)::`offset`
  must be a multiple of this value.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI-sType-sType"
  id="VUID-VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI-sType-sType"></a>
  VUID-VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CLUSTER_CULLING_SHADER_PROPERTIES_HUAWEI`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_HUAWEI_cluster_culling_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_HUAWEI_cluster_culling_shader.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
