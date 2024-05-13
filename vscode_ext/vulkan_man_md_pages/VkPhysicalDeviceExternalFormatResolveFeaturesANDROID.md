# VkPhysicalDeviceExternalFormatResolveFeaturesANDROID(3) Manual Page

## Name

VkPhysicalDeviceExternalFormatResolveFeaturesANDROID - Structure
describing whether external format resolves are supported



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceExternalFormatResolveFeaturesANDROID` structure is
defined as:

``` c
// Provided by VK_ANDROID_external_format_resolve
typedef struct VkPhysicalDeviceExternalFormatResolveFeaturesANDROID {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           externalFormatResolve;
} VkPhysicalDeviceExternalFormatResolveFeaturesANDROID;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-externalFormatResolve"></span>
  `externalFormatResolve` specifies whether external format resolves are
  supported.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceExternalFormatResolveFeaturesANDROID` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceExternalFormatResolveFeaturesANDROID` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceExternalFormatResolveFeaturesANDROID-sType-sType"
  id="VUID-VkPhysicalDeviceExternalFormatResolveFeaturesANDROID-sType-sType"></a>
  VUID-VkPhysicalDeviceExternalFormatResolveFeaturesANDROID-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FORMAT_RESOLVE_FEATURES_ANDROID`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ANDROID_external_format_resolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_format_resolve.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceExternalFormatResolveFeaturesANDROID"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
