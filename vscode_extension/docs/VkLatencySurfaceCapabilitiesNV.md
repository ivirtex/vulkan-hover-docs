# VkLatencySurfaceCapabilitiesNV(3) Manual Page

## Name

VkLatencySurfaceCapabilitiesNV - Structure describing surface optimized presentation modes for use with low latency mode



## [](#_c_specification)C Specification

The `VkLatencySurfaceCapabilitiesNV` structure is defined as:

```c++
// Provided by VK_NV_low_latency2
typedef struct VkLatencySurfaceCapabilitiesNV {
    VkStructureType      sType;
    const void*          pNext;
    uint32_t             presentModeCount;
    VkPresentModeKHR*    pPresentModes;
} VkLatencySurfaceCapabilitiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentModeCount` is the number of presentation modes provided.
- `pPresentModes` is list of presentation modes optimized for use with low latency mode with `presentModeCount` entries.

## [](#_description)Description

If `pPresentModes` is `NULL`, then the number of present modes that are optimized for use with low latency mode returned in `presentModeCount`. Otherwise, `presentModeCount` **must** be set by the application to the number of elements in the `pPresentModes` array, and on return the variable is overwritten with the number of values actually written to `pPresentModes`. If the value of `presentModeCount` is less than the number of optimized present modes, at most `presentModeCount` values will be written to `pPresentModes`.

Valid Usage (Implicit)

- [](#VUID-VkLatencySurfaceCapabilitiesNV-sType-sType)VUID-VkLatencySurfaceCapabilitiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_LATENCY_SURFACE_CAPABILITIES_NV`
- [](#VUID-VkLatencySurfaceCapabilitiesNV-pPresentModes-parameter)VUID-VkLatencySurfaceCapabilitiesNV-pPresentModes-parameter  
  If `presentModeCount` is not `0`, and `pPresentModes` is not `NULL`, `pPresentModes` **must** be a valid pointer to an array of `presentModeCount` [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) values

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkLatencySurfaceCapabilitiesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0