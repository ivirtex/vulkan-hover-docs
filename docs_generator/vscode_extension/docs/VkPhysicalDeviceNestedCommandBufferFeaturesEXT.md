# VkPhysicalDeviceNestedCommandBufferFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceNestedCommandBufferFeaturesEXT - Structure describing whether nested command buffers are supported by the implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceNestedCommandBufferFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_nested_command_buffer
typedef struct VkPhysicalDeviceNestedCommandBufferFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           nestedCommandBuffer;
    VkBool32           nestedCommandBufferRendering;
    VkBool32           nestedCommandBufferSimultaneousUse;
} VkPhysicalDeviceNestedCommandBufferFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- []()`nestedCommandBuffer` indicates the implementation supports nested command buffers, which allows [Secondary Command Buffers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary) to execute other [Secondary Command Buffers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary).
- []()`nestedCommandBufferRendering` indicates that it is valid to call [vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdExecuteCommands.html) inside a [Secondary Command Buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary) recorded with `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`.
- []()`nestedCommandBufferSimultaneousUse` indicates that the implementation supports nested command buffers with command buffers that are recorded with `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`.

## [](#_description)Description

If the `VkPhysicalDeviceNestedCommandBufferFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceNestedCommandBufferFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceNestedCommandBufferFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceNestedCommandBufferFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_NESTED_COMMAND_BUFFER_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_nested\_command\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_nested_command_buffer.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceNestedCommandBufferFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0