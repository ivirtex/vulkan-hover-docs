# VkPhysicalDeviceExtendedDynamicState3PropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceExtendedDynamicState3PropertiesEXT - Structure describing capabilities of extended dynamic state



## [](#_c_specification)C Specification

The `VkPhysicalDeviceExtendedDynamicState3PropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_extended_dynamic_state3
typedef struct VkPhysicalDeviceExtendedDynamicState3PropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           dynamicPrimitiveTopologyUnrestricted;
} VkPhysicalDeviceExtendedDynamicState3PropertiesEXT;
```

## [](#_members)Members

- []()`dynamicPrimitiveTopologyUnrestricted` indicates that the implementation allows `vkCmdSetPrimitiveTopology` to use a different [primitive topology class](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-primitive-topology-class) to the one specified in the active graphics pipeline.

## [](#_description)Description

If the `VkPhysicalDeviceExtendedDynamicState3PropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceExtendedDynamicState3PropertiesEXT-sType-sType)VUID-VkPhysicalDeviceExtendedDynamicState3PropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_3_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceExtendedDynamicState3PropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0