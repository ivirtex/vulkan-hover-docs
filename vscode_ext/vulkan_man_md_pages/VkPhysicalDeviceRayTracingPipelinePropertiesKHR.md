# VkPhysicalDeviceRayTracingPipelinePropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceRayTracingPipelinePropertiesKHR - Properties of the
physical device for ray tracing



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRayTracingPipelinePropertiesKHR` structure is
defined as:

``` c
// Provided by VK_KHR_ray_tracing_pipeline
typedef struct VkPhysicalDeviceRayTracingPipelinePropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           shaderGroupHandleSize;
    uint32_t           maxRayRecursionDepth;
    uint32_t           maxShaderGroupStride;
    uint32_t           shaderGroupBaseAlignment;
    uint32_t           shaderGroupHandleCaptureReplaySize;
    uint32_t           maxRayDispatchInvocationCount;
    uint32_t           shaderGroupHandleAlignment;
    uint32_t           maxRayHitAttributeSize;
} VkPhysicalDeviceRayTracingPipelinePropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `shaderGroupHandleSize` is the size in bytes of the shader header.

- <span id="limits-maxRayRecursionDepth"></span> `maxRayRecursionDepth`
  is the maximum number of levels of ray recursion allowed in a trace
  command.

- `maxShaderGroupStride` is the maximum stride in bytes allowed between
  shader groups in the shader binding table.

- `shaderGroupBaseAlignment` is the **required** alignment in bytes for
  the base of the shader binding table.

- `shaderGroupHandleCaptureReplaySize` is the number of bytes for the
  information required to do capture and replay for shader group
  handles.

- `maxRayDispatchInvocationCount` is the maximum number of ray
  generation shader invocations which **may** be produced by a single
  [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirectKHR.html) or
  [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html) command.

- `shaderGroupHandleAlignment` is the **required** alignment in bytes
  for each entry in a shader binding table. The value **must** be a
  power of two.

- `maxRayHitAttributeSize` is the maximum size in bytes for a ray
  attribute structure

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceRayTracingPipelinePropertiesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Limits specified by this structure **must** match those specified with
the same name in
[VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html).

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceRayTracingPipelinePropertiesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceRayTracingPipelinePropertiesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceRayTracingPipelinePropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PIPELINE_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRayTracingPipelinePropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
