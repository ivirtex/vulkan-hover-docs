# VkHostImageCopyDevicePerformanceQuery(3) Manual Page

## Name

VkHostImageCopyDevicePerformanceQuery - Struct containing information about optimality of device access



## [](#_c_specification)C Specification

To query if using `VK_IMAGE_USAGE_HOST_TRANSFER_BIT` has a negative impact on device performance when accessing an image, add `VK_IMAGE_USAGE_HOST_TRANSFER_BIT` to [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`usage`, and add a `VkHostImageCopyDevicePerformanceQuery` structure to the `pNext` chain of a [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html) structure passed to [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html). This structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkHostImageCopyDevicePerformanceQuery {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           optimalDeviceAccess;
    VkBool32           identicalMemoryLayout;
} VkHostImageCopyDevicePerformanceQuery;
```

or the equivalent

```c++
// Provided by VK_EXT_host_image_copy
typedef VkHostImageCopyDevicePerformanceQuery VkHostImageCopyDevicePerformanceQueryEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `optimalDeviceAccess` returns `VK_TRUE` if use of host image copy has no adverse effect on device access performance, compared to an image that is created with exact same creation parameters, and bound to the same [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), except that `VK_IMAGE_USAGE_HOST_TRANSFER_BIT` is replaced with `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` and `VK_IMAGE_USAGE_TRANSFER_DST_BIT`.
- `identicalMemoryLayout` returns `VK_TRUE` if use of host image copy has no impact on memory layout compared to an image that is created with exact same creation parameters, and bound to the same [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), except that `VK_IMAGE_USAGE_HOST_TRANSFER_BIT` is replaced with `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` and `VK_IMAGE_USAGE_TRANSFER_DST_BIT`.

## [](#_description)Description

The implementation **may** return `VK_FALSE` in `optimalDeviceAccess` if `identicalMemoryLayout` is `VK_FALSE`. If `identicalMemoryLayout` is `VK_TRUE`, `optimalDeviceAccess` **must** be `VK_TRUE`.

The implementation **may** return `VK_TRUE` in `optimalDeviceAccess` while `identicalMemoryLayout` is `VK_FALSE`. In this situation, any device performance impact **should** not be measurable.

If [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`format` is a block-compressed format and [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html) returns `VK_SUCCESS`, the implementation **must** return `VK_TRUE` in `optimalDeviceAccess`.

Note

Applications can make use of `optimalDeviceAccess` to determine their resource copying strategy. If a resource is expected to be accessed more on device than on the host, and the implementation considers the resource sub-optimally accessed, it is likely better to use device copies instead.

Note

Layout not being identical yet still considered optimal for device access could happen if the implementation has different memory layout patterns, some of which are easier to access on the host.

Note

The most practical reason for `optimalDeviceAccess` to be `VK_FALSE` is that host image access may disable framebuffer compression where it would otherwise have been enabled. This represents far more efficient host image access since no compression algorithm is required to read or write to the image, but it would impact device access performance. Some implementations may only set `optimalDeviceAccess` to `VK_FALSE` if certain conditions are met, such as specific image usage flags or creation flags.

Valid Usage (Implicit)

- [](#VUID-VkHostImageCopyDevicePerformanceQuery-sType-sType)VUID-VkHostImageCopyDevicePerformanceQuery-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_HOST_IMAGE_COPY_DEVICE_PERFORMANCE_QUERY`

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkHostImageCopyDevicePerformanceQuery)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0