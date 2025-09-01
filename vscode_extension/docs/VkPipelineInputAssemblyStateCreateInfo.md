# VkPipelineInputAssemblyStateCreateInfo(3) Manual Page

## Name

VkPipelineInputAssemblyStateCreateInfo - Structure specifying parameters of a newly created pipeline input assembly state



## [](#_c_specification)C Specification

Drawing can be achieved in two modes:

- [Programmable Mesh Shading](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-mesh-shading), the mesh shader assembles primitives, or
- [Programmable Primitive Shading](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-primitive-shading), the input primitives are assembled as follows.

Each draw is made up of zero or more vertices and zero or more instances, which are processed by the device and result in the assembly of primitives. Primitives are assembled according to the `pInputAssemblyState` member of the [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html) structure, which is of type `VkPipelineInputAssemblyStateCreateInfo`:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineInputAssemblyStateCreateInfo {
    VkStructureType                            sType;
    const void*                                pNext;
    VkPipelineInputAssemblyStateCreateFlags    flags;
    VkPrimitiveTopology                        topology;
    VkBool32                                   primitiveRestartEnable;
} VkPipelineInputAssemblyStateCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `topology` is a [VkPrimitiveTopology](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrimitiveTopology.html) defining the primitive topology, as described below.
- `primitiveRestartEnable` controls whether a special vertex index value is treated as restarting the assembly of primitives. This enable only applies to indexed draws ([vkCmdDrawIndexed](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexed.html), [vkCmdDrawMultiIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMultiIndexedEXT.html), and [vkCmdDrawIndexedIndirect](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexedIndirect.html)), and the special index value is either 0xFFFFFFFF when the `indexType` parameter of Vulkan 1.4 or `vkCmdBindIndexBuffer2` or `vkCmdBindIndexBuffer` is equal to `VK_INDEX_TYPE_UINT32`; 0xFF when `indexType` is equal to `VK_INDEX_TYPE_UINT8`; or 0xFFFF when `indexType` is equal to `VK_INDEX_TYPE_UINT16`. Primitive restart is not allowed for “list” topologies, unless one of the features [`primitiveTopologyPatchListRestart`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-primitiveTopologyPatchListRestart) (for `VK_PRIMITIVE_TOPOLOGY_PATCH_LIST`) or [`primitiveTopologyListRestart`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-primitiveTopologyListRestart) (for all other list topologies) is enabled.

## [](#_description)Description

Restarting the assembly of primitives discards the most recent index values if those elements formed an incomplete primitive, and restarts the primitive assembly using the subsequent indices, but only assembling the immediately following element through the end of the originally specified elements. The primitive restart index value comparison is performed before adding the `vertexOffset` value to the index value.

Valid Usage

- [](#VUID-VkPipelineInputAssemblyStateCreateInfo-topology-06252)VUID-VkPipelineInputAssemblyStateCreateInfo-topology-06252  
  If the [`primitiveTopologyListRestart`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-primitiveTopologyListRestart) feature is not enabled, and `topology` is `VK_PRIMITIVE_TOPOLOGY_POINT_LIST`, `VK_PRIMITIVE_TOPOLOGY_LINE_LIST`, `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST`, `VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY`, or `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY`, `primitiveRestartEnable` **must** be `VK_FALSE`
- [](#VUID-VkPipelineInputAssemblyStateCreateInfo-topology-06253)VUID-VkPipelineInputAssemblyStateCreateInfo-topology-06253  
  If the [`primitiveTopologyPatchListRestart`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-primitiveTopologyPatchListRestart) feature is not enabled, and `topology` is `VK_PRIMITIVE_TOPOLOGY_PATCH_LIST`, `primitiveRestartEnable` **must** be `VK_FALSE`
- [](#VUID-VkPipelineInputAssemblyStateCreateInfo-topology-00429)VUID-VkPipelineInputAssemblyStateCreateInfo-topology-00429  
  If the [`geometryShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-geometryShader) feature is not enabled, `topology` **must** not be any of `VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY`, `VK_PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY`, `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY` or `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY`
- [](#VUID-VkPipelineInputAssemblyStateCreateInfo-topology-00430)VUID-VkPipelineInputAssemblyStateCreateInfo-topology-00430  
  If the [`tessellationShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tessellationShader) feature is not enabled, `topology` **must** not be `VK_PRIMITIVE_TOPOLOGY_PATCH_LIST`
- [](#VUID-VkPipelineInputAssemblyStateCreateInfo-triangleFans-04452)VUID-VkPipelineInputAssemblyStateCreateInfo-triangleFans-04452  
  If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`triangleFans` is `VK_FALSE`, `topology` **must** not be `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_FAN`

Valid Usage (Implicit)

- [](#VUID-VkPipelineInputAssemblyStateCreateInfo-sType-sType)VUID-VkPipelineInputAssemblyStateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO`
- [](#VUID-VkPipelineInputAssemblyStateCreateInfo-pNext-pNext)VUID-VkPipelineInputAssemblyStateCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPipelineInputAssemblyStateCreateInfo-flags-zerobitmask)VUID-VkPipelineInputAssemblyStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkPipelineInputAssemblyStateCreateInfo-topology-parameter)VUID-VkPipelineInputAssemblyStateCreateInfo-topology-parameter  
  `topology` **must** be a valid [VkPrimitiveTopology](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrimitiveTopology.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkPipelineInputAssemblyStateCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineInputAssemblyStateCreateFlags.html), [VkPrimitiveTopology](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrimitiveTopology.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineInputAssemblyStateCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0