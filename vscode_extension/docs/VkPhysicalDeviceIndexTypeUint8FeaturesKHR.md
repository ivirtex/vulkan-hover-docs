# VkPhysicalDeviceIndexTypeUint8FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceIndexTypeUint8FeaturesKHR - Structure describing whether
uint8 index type can be used



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceIndexTypeUint8FeaturesKHR` structure is defined as:

``` c
// Provided by VK_KHR_index_type_uint8
typedef struct VkPhysicalDeviceIndexTypeUint8FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           indexTypeUint8;
} VkPhysicalDeviceIndexTypeUint8FeaturesKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_index_type_uint8
typedef VkPhysicalDeviceIndexTypeUint8FeaturesKHR VkPhysicalDeviceIndexTypeUint8FeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-indexTypeUint8"></span> `indexTypeUint8` indicates
  that `VK_INDEX_TYPE_UINT8_KHR` can be used with
  [vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer2KHR.html) and
  [vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer.html).

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceIndexTypeUint8FeaturesKHR` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceIndexTypeUint8FeaturesKHR` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceIndexTypeUint8FeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceIndexTypeUint8FeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceIndexTypeUint8FeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INDEX_TYPE_UINT8_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_index_type_uint8](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_index_type_uint8.html),
[VK_KHR_index_type_uint8](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_index_type_uint8.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceIndexTypeUint8FeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
