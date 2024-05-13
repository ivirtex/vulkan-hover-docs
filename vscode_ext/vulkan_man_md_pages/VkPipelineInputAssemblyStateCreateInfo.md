# VkPipelineInputAssemblyStateCreateInfo(3) Manual Page

## Name

VkPipelineInputAssemblyStateCreateInfo - Structure specifying parameters
of a newly created pipeline input assembly state



## <a href="#_c_specification" class="anchor"></a>C Specification

Drawing can be achieved in two modes:

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-mesh-shading"
  target="_blank" rel="noopener">Programmable Mesh Shading</a>, the mesh
  shader assembles primitives, or

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-primitive-shading"
  target="_blank" rel="noopener">Programmable Primitive Shading</a>, the
  input primitives are assembled as follows.

Each draw is made up of zero or more vertices and zero or more
instances, which are processed by the device and result in the assembly
of primitives. Primitives are assembled according to the
`pInputAssemblyState` member of the
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)
structure, which is of type `VkPipelineInputAssemblyStateCreateInfo`:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkPipelineInputAssemblyStateCreateInfo {
    VkStructureType                            sType;
    const void*                                pNext;
    VkPipelineInputAssemblyStateCreateFlags    flags;
    VkPrimitiveTopology                        topology;
    VkBool32                                   primitiveRestartEnable;
} VkPipelineInputAssemblyStateCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `topology` is a [VkPrimitiveTopology](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html)
  defining the primitive topology, as described below.

- `primitiveRestartEnable` controls whether a special vertex index value
  is treated as restarting the assembly of primitives. This enable only
  applies to indexed draws ([vkCmdDrawIndexed](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexed.html),
  [vkCmdDrawMultiIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMultiIndexedEXT.html), and
  [vkCmdDrawIndexedIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirect.html)), and the
  special index value is either 0xFFFFFFFF when the `indexType`
  parameter of `vkCmdBindIndexBuffer2KHR` or `vkCmdBindIndexBuffer` is
  equal to `VK_INDEX_TYPE_UINT32`, 0xFF when `indexType` is equal to
  `VK_INDEX_TYPE_UINT8_KHR`, or 0xFFFF when `indexType` is equal to
  `VK_INDEX_TYPE_UINT16`. Primitive restart is not allowed for “list”
  topologies, unless one of the features <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitiveTopologyPatchListRestart"
  target="_blank"
  rel="noopener"><code>primitiveTopologyPatchListRestart</code></a> (for
  `VK_PRIMITIVE_TOPOLOGY_PATCH_LIST`) or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitiveTopologyListRestart"
  target="_blank"
  rel="noopener"><code>primitiveTopologyListRestart</code></a> (for all
  other list topologies) is enabled.

## <a href="#_description" class="anchor"></a>Description

Restarting the assembly of primitives discards the most recent index
values if those elements formed an incomplete primitive, and restarts
the primitive assembly using the subsequent indices, but only assembling
the immediately following element through the end of the originally
specified elements. The primitive restart index value comparison is
performed before adding the `vertexOffset` value to the index value.

Valid Usage

- <a href="#VUID-VkPipelineInputAssemblyStateCreateInfo-topology-06252"
  id="VUID-VkPipelineInputAssemblyStateCreateInfo-topology-06252"></a>
  VUID-VkPipelineInputAssemblyStateCreateInfo-topology-06252  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitiveTopologyListRestart"
  target="_blank"
  rel="noopener"><code>primitiveTopologyListRestart</code></a> feature
  is not enabled, and `topology` is `VK_PRIMITIVE_TOPOLOGY_POINT_LIST`,
  `VK_PRIMITIVE_TOPOLOGY_LINE_LIST`,
  `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST`,
  `VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY`, or
  `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY`,
  `primitiveRestartEnable` **must** be `VK_FALSE`

- <a href="#VUID-VkPipelineInputAssemblyStateCreateInfo-topology-06253"
  id="VUID-VkPipelineInputAssemblyStateCreateInfo-topology-06253"></a>
  VUID-VkPipelineInputAssemblyStateCreateInfo-topology-06253  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitiveTopologyPatchListRestart"
  target="_blank"
  rel="noopener"><code>primitiveTopologyPatchListRestart</code></a>
  feature is not enabled, and `topology` is
  `VK_PRIMITIVE_TOPOLOGY_PATCH_LIST`, `primitiveRestartEnable` **must**
  be `VK_FALSE`

- <a href="#VUID-VkPipelineInputAssemblyStateCreateInfo-topology-00429"
  id="VUID-VkPipelineInputAssemblyStateCreateInfo-topology-00429"></a>
  VUID-VkPipelineInputAssemblyStateCreateInfo-topology-00429  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-geometryShader"
  target="_blank" rel="noopener"><code>geometryShader</code></a> feature
  is not enabled, `topology` **must** not be any of
  `VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY`,
  `VK_PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY`,
  `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY` or
  `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY`

- <a href="#VUID-VkPipelineInputAssemblyStateCreateInfo-topology-00430"
  id="VUID-VkPipelineInputAssemblyStateCreateInfo-topology-00430"></a>
  VUID-VkPipelineInputAssemblyStateCreateInfo-topology-00430  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-tessellationShader"
  target="_blank" rel="noopener"><code>tessellationShader</code></a>
  feature is not enabled, `topology` **must** not be
  `VK_PRIMITIVE_TOPOLOGY_PATCH_LIST`

- <a
  href="#VUID-VkPipelineInputAssemblyStateCreateInfo-triangleFans-04452"
  id="VUID-VkPipelineInputAssemblyStateCreateInfo-triangleFans-04452"></a>
  VUID-VkPipelineInputAssemblyStateCreateInfo-triangleFans-04452  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`triangleFans`
  is `VK_FALSE`, `topology` **must** not be
  `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_FAN`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineInputAssemblyStateCreateInfo-sType-sType"
  id="VUID-VkPipelineInputAssemblyStateCreateInfo-sType-sType"></a>
  VUID-VkPipelineInputAssemblyStateCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO`

- <a href="#VUID-VkPipelineInputAssemblyStateCreateInfo-pNext-pNext"
  id="VUID-VkPipelineInputAssemblyStateCreateInfo-pNext-pNext"></a>
  VUID-VkPipelineInputAssemblyStateCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkPipelineInputAssemblyStateCreateInfo-flags-zerobitmask"
  id="VUID-VkPipelineInputAssemblyStateCreateInfo-flags-zerobitmask"></a>
  VUID-VkPipelineInputAssemblyStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkPipelineInputAssemblyStateCreateInfo-topology-parameter"
  id="VUID-VkPipelineInputAssemblyStateCreateInfo-topology-parameter"></a>
  VUID-VkPipelineInputAssemblyStateCreateInfo-topology-parameter  
  `topology` **must** be a valid
  [VkPrimitiveTopology](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkPipelineInputAssemblyStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInputAssemblyStateCreateFlags.html),
[VkPrimitiveTopology](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineInputAssemblyStateCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
