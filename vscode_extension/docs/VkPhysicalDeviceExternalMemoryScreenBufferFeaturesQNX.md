# VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX(3) Manual Page

## Name

VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX - Structure describing QNX Screen Buffer features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX` structure is defined as:

```c++
// Provided by VK_QNX_external_memory_screen_buffer
typedef struct VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           screenBufferImport;
} VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX` structure describe the following features:

## [](#_description)Description

- []()`screenBufferImport` indicates whether QNX Screen buffer import functionality is supported. If `screenBufferImport` is `VK_TRUE`, [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) supports importing `_screen_buffer` from applications. In this case, the application is responsible for the resource management of the `_screen_buffer`.

Table 1. Functionality Supported for QNX Screen Buffer Features

Features

Functionality

`screenBufferImport`

[VkImportScreenBufferInfoQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportScreenBufferInfoQNX.html)

Always supported1

[vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetScreenBufferPropertiesQNX.html), [VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferPropertiesQNX.html), [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferFormatPropertiesQNX.html), [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatQNX.html)

1

Functionality in this row is always available.

The [Functionality supported for QNX Screen buffer features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalscreenbuffer-table) table summarizes the functionality enabled by the `VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX` structure. Each entry in the body of the table summarizes the functionality that **can** be used when the given features are supported and enabled. This summarizes Valid Usage statements that are added elsewhere in this specification.

If the `VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX-sType-sType)VUID-VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_SCREEN_BUFFER_FEATURES_QNX`

## [](#_see_also)See Also

[VK\_QNX\_external\_memory\_screen\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QNX_external_memory_screen_buffer.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0