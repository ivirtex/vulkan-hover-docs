# VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV - Structure
describing feature to control ray tracing invocation reordering



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV` structure is
defined as:

``` c
// Provided by VK_NV_ray_tracing_invocation_reorder
typedef struct VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           rayTracingInvocationReorder;
} VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-rayTracingInvocationReorder"></span>
  `rayTracingInvocationReorder` indicates that the implementation
  supports `SPV_NV_shader_invocation_reorder`.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_INVOCATION_REORDER_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing_invocation_reorder](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing_invocation_reorder.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
