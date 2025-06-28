# VkPhysicalDeviceToolProperties(3) Manual Page

## Name

VkPhysicalDeviceToolProperties - Structure providing information about an active tool



## [](#_c_specification)C Specification

The [VkPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceToolProperties.html) structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceToolProperties {
    VkStructureType       sType;
    void*                 pNext;
    char                  name[VK_MAX_EXTENSION_NAME_SIZE];
    char                  version[VK_MAX_EXTENSION_NAME_SIZE];
    VkToolPurposeFlags    purposes;
    char                  description[VK_MAX_DESCRIPTION_SIZE];
    char                  layer[VK_MAX_EXTENSION_NAME_SIZE];
} VkPhysicalDeviceToolProperties;
```

or the equivalent

```c++
// Provided by VK_EXT_tooling_info
typedef VkPhysicalDeviceToolProperties VkPhysicalDeviceToolPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `name` is a null-terminated UTF-8 string containing the name of the tool.
- `version` is a null-terminated UTF-8 string containing the version of the tool.
- `purposes` is a bitmask of [VkToolPurposeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkToolPurposeFlagBits.html) which is populated with purposes supported by the tool.
- `description` is a null-terminated UTF-8 string containing a description of the tool.
- `layer` is a null-terminated UTF-8 string containing the name of the layer implementing the tool, if the tool is implemented in a layer - otherwise it **may** be an empty string.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceToolProperties-sType-sType)VUID-VkPhysicalDeviceToolProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TOOL_PROPERTIES`
- [](#VUID-VkPhysicalDeviceToolProperties-pNext-pNext)VUID-VkPhysicalDeviceToolProperties-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_EXT\_tooling\_info](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_tooling_info.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkToolPurposeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkToolPurposeFlags.html), [vkGetPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceToolProperties.html), [vkGetPhysicalDeviceToolPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceToolPropertiesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceToolProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0