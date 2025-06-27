# VkImageViewUsageCreateInfo(3) Manual Page

## Name

VkImageViewUsageCreateInfo - Specify the intended usage of an image view



## <a href="#_c_specification" class="anchor"></a>C Specification

The set of usages for the created image view **can** be restricted
compared to the parent image’s `usage` flags by adding a
`VkImageViewUsageCreateInfo` structure to the `pNext` chain of
[VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html).

The `VkImageViewUsageCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkImageViewUsageCreateInfo {
    VkStructureType      sType;
    const void*          pNext;
    VkImageUsageFlags    usage;
} VkImageViewUsageCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_maintenance2
typedef VkImageViewUsageCreateInfo VkImageViewUsageCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `usage` is a bitmask of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) specifying allowed
  usages of the image view.

## <a href="#_description" class="anchor"></a>Description

When this structure is chained to
[VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html) the `usage` field
overrides the implicit `usage` parameter inherited from image creation
time and its value is used instead for the purposes of determining the
valid usage conditions of
[VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html).

Valid Usage (Implicit)

- <a href="#VUID-VkImageViewUsageCreateInfo-sType-sType"
  id="VUID-VkImageViewUsageCreateInfo-sType-sType"></a>
  VUID-VkImageViewUsageCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO`

- <a href="#VUID-VkImageViewUsageCreateInfo-usage-parameter"
  id="VUID-VkImageViewUsageCreateInfo-usage-parameter"></a>
  VUID-VkImageViewUsageCreateInfo-usage-parameter  
  `usage` **must** be a valid combination of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) values

- <a href="#VUID-VkImageViewUsageCreateInfo-usage-requiredbitmask"
  id="VUID-VkImageViewUsageCreateInfo-usage-requiredbitmask"></a>
  VUID-VkImageViewUsageCreateInfo-usage-requiredbitmask  
  `usage` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageViewUsageCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
