# VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT - Structure describing
whether querying and providing an identifier of a shader module is
supported by the implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_shader_module_identifier
typedef struct VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderModuleIdentifier;
} VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-shaderModuleIdentifier"></span>
  `shaderModuleIdentifier` indicates whether the implementation supports
  querying an identifier of a [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) or
  [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) structure,
  and creating pipelines from identifiers only.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_MODULE_IDENTIFIER_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_module_identifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_module_identifier.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
