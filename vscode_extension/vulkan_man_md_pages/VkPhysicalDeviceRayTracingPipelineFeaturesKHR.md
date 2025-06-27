# VkPhysicalDeviceRayTracingPipelineFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceRayTracingPipelineFeaturesKHR - Structure describing the
ray tracing features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRayTracingPipelineFeaturesKHR` structure is defined
as:

``` c
// Provided by VK_KHR_ray_tracing_pipeline
typedef struct VkPhysicalDeviceRayTracingPipelineFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           rayTracingPipeline;
    VkBool32           rayTracingPipelineShaderGroupHandleCaptureReplay;
    VkBool32           rayTracingPipelineShaderGroupHandleCaptureReplayMixed;
    VkBool32           rayTracingPipelineTraceRaysIndirect;
    VkBool32           rayTraversalPrimitiveCulling;
} VkPhysicalDeviceRayTracingPipelineFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-rayTracingPipeline"></span> `rayTracingPipeline`
  indicates whether the implementation supports the ray tracing pipeline
  functionality. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-tracing"
  target="_blank" rel="noopener">Ray Tracing</a>.

- <span id="features-rayTracingPipelineShaderGroupHandleCaptureReplay"></span>
  `rayTracingPipelineShaderGroupHandleCaptureReplay` indicates whether
  the implementation supports saving and reusing shader group handles,
  e.g. for trace capture and replay.

- <span id="features-rayTracingPipelineShaderGroupHandleCaptureReplayMixed"></span>
  `rayTracingPipelineShaderGroupHandleCaptureReplayMixed` indicates
  whether the implementation supports reuse of shader group handles
  being arbitrarily mixed with creation of non-reused shader group
  handles. If this is `VK_FALSE`, all reused shader group handles
  **must** be specified before any non-reused handles **may** be
  created.

- <span id="features-rayTracingPipelineTraceRaysIndirect"></span>
  `rayTracingPipelineTraceRaysIndirect` indicates whether the
  implementation supports indirect ray tracing commands, e.g.
  [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirectKHR.html).

- <span id="features-rayTraversalPrimitiveCulling"></span>
  `rayTraversalPrimitiveCulling` indicates whether the implementation
  supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-traversal-culling-primitive"
  target="_blank" rel="noopener">primitive culling during ray
  traversal</a>.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceRayTracingPipelineFeaturesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceRayTracingPipelineFeaturesKHR` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage

- <a
  href="#VUID-VkPhysicalDeviceRayTracingPipelineFeaturesKHR-rayTracingPipelineShaderGroupHandleCaptureReplayMixed-03575"
  id="VUID-VkPhysicalDeviceRayTracingPipelineFeaturesKHR-rayTracingPipelineShaderGroupHandleCaptureReplayMixed-03575"></a>
  VUID-VkPhysicalDeviceRayTracingPipelineFeaturesKHR-rayTracingPipelineShaderGroupHandleCaptureReplayMixed-03575  
  If `rayTracingPipelineShaderGroupHandleCaptureReplayMixed` is
  `VK_TRUE`, `rayTracingPipelineShaderGroupHandleCaptureReplay` **must**
  also be `VK_TRUE`

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceRayTracingPipelineFeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceRayTracingPipelineFeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceRayTracingPipelineFeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PIPELINE_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRayTracingPipelineFeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
