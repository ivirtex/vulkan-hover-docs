# VkPhysicalDeviceLayeredApiPropertiesListKHR(3) Manual Page

## Name

VkPhysicalDeviceLayeredApiPropertiesListKHR - Structure describing
layered implementations underneath the Vulkan physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceLayeredApiPropertiesListKHR` structure is defined
as:

``` c
// Provided by VK_KHR_maintenance7
typedef struct VkPhysicalDeviceLayeredApiPropertiesListKHR {
    VkStructureType                             sType;
    void*                                       pNext;
    uint32_t                                    layeredApiCount;
    VkPhysicalDeviceLayeredApiPropertiesKHR*    pLayeredApis;
} VkPhysicalDeviceLayeredApiPropertiesListKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `layeredApiCount` is an integer related to the number of layered
  implementations underneath the Vulkan physical device, as described
  below.

- `pLayeredApis` is a pointer to an array of
  [VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html)
  in which information regarding the layered implementations underneath
  the Vulkan physical device are returned.

## <a href="#_description" class="anchor"></a>Description

If `pLayeredApis` is `NULL`, then the number of layered implementations
that are underneath the top-most Vulkan physical device (i.e. the one
returned by
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html))
is returned in `layeredApiCount`. Otherwise, `layeredApiCount` **must**
be set by the application to the number of elements in the
`pLayeredApis` array, and on return the variable is overwritten with the
number of values actually written to `pLayeredApis`. If the value of
`layeredApiCount` is less than the number of layered implementations
underneath the Vulkan physical device, at most `layeredApiCount` values
will be written to `pLayeredApis`. An implementation that is not a layer
will return 0 in `layeredApiCount`.

In the presence of multiple layered implementations, each element of
`pLayeredApis` corresponds to an API implementation that is implemented
on top of the API at the previous index. If there are layered
implementations underneath a non-Vulkan implementation, they may not be
visible in this query as the corresponding APIs may lack such a query.

If the `VkPhysicalDeviceLayeredApiPropertiesListKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceLayeredApiPropertiesListKHR-sType-sType"
  id="VUID-VkPhysicalDeviceLayeredApiPropertiesListKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceLayeredApiPropertiesListKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_API_PROPERTIES_LIST_KHR`

- <a
  href="#VUID-VkPhysicalDeviceLayeredApiPropertiesListKHR-pLayeredApis-parameter"
  id="VUID-VkPhysicalDeviceLayeredApiPropertiesListKHR-pLayeredApis-parameter"></a>
  VUID-VkPhysicalDeviceLayeredApiPropertiesListKHR-pLayeredApis-parameter  
  If `layeredApiCount` is not `0`, and `pLayeredApis` is not `NULL`,
  `pLayeredApis` **must** be a valid pointer to an array of
  `layeredApiCount`
  [VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance7](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance7.html),
[VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceLayeredApiPropertiesListKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
