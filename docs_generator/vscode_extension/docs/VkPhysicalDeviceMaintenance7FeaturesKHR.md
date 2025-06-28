# VkPhysicalDeviceMaintenance7FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceMaintenance7FeaturesKHR - Structure describing whether the implementation supports maintenance7 functionality



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMaintenance7FeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_maintenance7
typedef struct VkPhysicalDeviceMaintenance7FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           maintenance7;
} VkPhysicalDeviceMaintenance7FeaturesKHR;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maintenance7` indicates that the implementation supports the following:
  
  - The `VK_RENDERING_CONTENTS_INLINE_BIT_KHR` and `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_KHR` flags **can** be used to record commands in render pass instances both inline and in secondary command buffers executed with [vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdExecuteCommands.html) for dynamic rendering and legacy render passes respectively.
  - Querying information regarding the underlying devices in environments where the Vulkan implementation is provided through layered implementations. This is done by chaining [VkPhysicalDeviceLayeredApiPropertiesListKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiPropertiesListKHR.html) to [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html).
  - New limits which indicate the maximum total count of dynamic uniform buffers and dynamic storage buffers that **can** be included in a pipeline layout.
  - 32-bit timestamp queries **must** wrap on overflow
  - A property that indicates whether a fragment shading rate attachment can have a size that is too small to cover a specified render area.
  - A property that indicates support for writing to one aspect of a depth/stencil attachment without performing a read-modify-write operation on the other aspect

## [](#_description)Description

If the `VkPhysicalDeviceMaintenance7FeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceMaintenance7FeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMaintenance7FeaturesKHR-sType-sType)VUID-VkPhysicalDeviceMaintenance7FeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_7_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_maintenance7](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance7.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMaintenance7FeaturesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0