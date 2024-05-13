# VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX(3) Manual Page

## Name

VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX - Structure
describing multiview limits that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX` structure
is defined as:

``` c
// Provided by VK_NVX_multiview_per_view_attributes
typedef struct VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           perViewPositionAllComponents;
} VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-perViewPositionAllComponents"></span>
  `perViewPositionAllComponents` is `VK_TRUE` if the implementation
  supports per-view position values that differ in components other than
  the X component.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX-sType-sType"
  id="VUID-VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX-sType-sType"></a>
  VUID-VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_ATTRIBUTES_PROPERTIES_NVX`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NVX_multiview_per_view_attributes](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_multiview_per_view_attributes.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
