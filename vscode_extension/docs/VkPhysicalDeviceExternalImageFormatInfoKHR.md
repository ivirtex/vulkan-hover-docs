# VkPhysicalDeviceExternalImageFormatInfo(3) Manual Page

## Name

VkPhysicalDeviceExternalImageFormatInfo - Structure specifying external image creation parameters



## [](#_c_specification)C Specification

To determine the image capabilities compatible with an external memory handle type, add a [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalImageFormatInfo.html) structure to the `pNext` chain of the [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html) structure and a `VkExternalImageFormatProperties` structure to the `pNext` chain of the [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html) structure.

The `VkPhysicalDeviceExternalImageFormatInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceExternalImageFormatInfo {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkPhysicalDeviceExternalImageFormatInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_external_memory_capabilities
typedef VkPhysicalDeviceExternalImageFormatInfo VkPhysicalDeviceExternalImageFormatInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the memory handle type that will be used with the memory associated with the image.

## [](#_description)Description

If `handleType` is 0, [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html) will behave as if [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalImageFormatInfo.html) was not present, and [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatProperties.html) will be ignored.

If `handleType` is not compatible with the `format`, `type`, `tiling`, `usage`, and `flags` specified in [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html), then [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html) returns `VK_ERROR_FORMAT_NOT_SUPPORTED`.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceExternalImageFormatInfo-sType-sType)VUID-VkPhysicalDeviceExternalImageFormatInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO`
- [](#VUID-VkPhysicalDeviceExternalImageFormatInfo-handleType-parameter)VUID-VkPhysicalDeviceExternalImageFormatInfo-handleType-parameter  
  If `handleType` is not `0`, `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceExternalImageFormatInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0