# VkAndroidHardwareBufferUsageANDROID(3) Manual Page

## Name

VkAndroidHardwareBufferUsageANDROID - Struct containing Android hardware buffer usage flags



## [](#_c_specification)C Specification

To obtain optimal Android hardware buffer usage flags for specific image creation parameters, add a `VkAndroidHardwareBufferUsageANDROID` structure to the `pNext` chain of a [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html) structure passed to [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html). This structure is defined as:

```c++
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
typedef struct VkAndroidHardwareBufferUsageANDROID {
    VkStructureType    sType;
    void*              pNext;
    uint64_t           androidHardwareBufferUsage;
} VkAndroidHardwareBufferUsageANDROID;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `androidHardwareBufferUsage` returns the Android hardware buffer usage flags.

## [](#_description)Description

The `androidHardwareBufferUsage` field **must** include Android hardware buffer usage flags listed in the [AHardwareBuffer Usage Equivalence](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-android-hardware-buffer-usage) table when the corresponding Vulkan image usage or image creation flags are included in the `usage` or `flags` fields of [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html). It **must** include at least one GPU usage flag (`AHARDWAREBUFFER_USAGE_GPU_*`), even if none of the corresponding Vulkan usages or flags are requested.

Note

Requiring at least one GPU usage flag ensures that Android hardware buffer memory will be allocated in a memory pool accessible to the Vulkan implementation, and that specializing the memory layout based on usage flags does not prevent it from being compatible with Vulkan. Implementations **may** avoid unnecessary restrictions caused by this requirement by using vendor usage flags to indicate that only the Vulkan uses indicated in [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html) are required.

Valid Usage (Implicit)

- [](#VUID-VkAndroidHardwareBufferUsageANDROID-sType-sType)VUID-VkAndroidHardwareBufferUsageANDROID-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_USAGE_ANDROID`

## [](#_see_also)See Also

[VK\_ANDROID\_external\_memory\_android\_hardware\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAndroidHardwareBufferUsageANDROID).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0