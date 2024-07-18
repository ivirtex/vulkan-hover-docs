# VkLatencySurfaceCapabilitiesNV(3) Manual Page

## Name

VkLatencySurfaceCapabilitiesNV - Structure describing surface optimized
presentation modes for use with low latency mode



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkLatencySurfaceCapabilitiesNV` structure is defined as:

``` c
// Provided by VK_NV_low_latency2
typedef struct VkLatencySurfaceCapabilitiesNV {
    VkStructureType      sType;
    const void*          pNext;
    uint32_t             presentModeCount;
    VkPresentModeKHR*    pPresentModes;
} VkLatencySurfaceCapabilitiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `presentModeCount` is the number of presentation modes provided.

- `pPresentModes` is list of presentation modes optimized for use with
  low latency mode with `presentModeCount` entries.

## <a href="#_description" class="anchor"></a>Description

If `pPresentModes` is `NULL`, then the number of present modes that are
optimized for use with low latency mode returned in `presentModeCount`.
Otherwise, `presentModeCount` must be set by the application to the
number of elements in the `pPresentModes` array, and on return the
variable is overwritten with the number of values actually written to
`pPresentModes`. If the value of `presentModeCount` is less than the
number of optimized present modes, at most `presentModeCount` values
will be written to `pPresentModes`.

Valid Usage (Implicit)

- <a href="#VUID-VkLatencySurfaceCapabilitiesNV-sType-sType"
  id="VUID-VkLatencySurfaceCapabilitiesNV-sType-sType"></a>
  VUID-VkLatencySurfaceCapabilitiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_LATENCY_SURFACE_CAPABILITIES_NV`

- <a href="#VUID-VkLatencySurfaceCapabilitiesNV-pPresentModes-parameter"
  id="VUID-VkLatencySurfaceCapabilitiesNV-pPresentModes-parameter"></a>
  VUID-VkLatencySurfaceCapabilitiesNV-pPresentModes-parameter  
  If `presentModeCount` is not `0`, and `pPresentModes` is not `NULL`,
  `pPresentModes` **must** be a valid pointer to an array of
  `presentModeCount` [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkLatencySurfaceCapabilitiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
