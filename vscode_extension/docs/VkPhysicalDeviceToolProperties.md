# VkPhysicalDeviceToolProperties(3) Manual Page

## Name

VkPhysicalDeviceToolProperties - Structure providing information about
an active tool



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceToolProperties.html)
structure is defined as:

``` c
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

``` c
// Provided by VK_EXT_tooling_info
typedef VkPhysicalDeviceToolProperties VkPhysicalDeviceToolPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `name` is a null-terminated UTF-8 string containing the name of the
  tool.

- `version` is a null-terminated UTF-8 string containing the version of
  the tool.

- `purposes` is a bitmask of
  [VkToolPurposeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkToolPurposeFlagBits.html) which is populated
  with purposes supported by the tool.

- `description` is a null-terminated UTF-8 string containing a
  description of the tool.

- `layer` is a null-terminated UTF-8 string containing the name of the
  layer implementing the tool, if the tool is implemented in a layer -
  otherwise it **may** be an empty string.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceToolProperties-sType-sType"
  id="VUID-VkPhysicalDeviceToolProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceToolProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TOOL_PROPERTIES`

- <a href="#VUID-VkPhysicalDeviceToolProperties-pNext-pNext"
  id="VUID-VkPhysicalDeviceToolProperties-pNext-pNext"></a>
  VUID-VkPhysicalDeviceToolProperties-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_tooling_info](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_tooling_info.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkToolPurposeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkToolPurposeFlags.html),
[vkGetPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceToolProperties.html),
[vkGetPhysicalDeviceToolPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceToolPropertiesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceToolProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
