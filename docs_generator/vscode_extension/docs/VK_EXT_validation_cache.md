# VK\_EXT\_validation\_cache(3) Manual Page

## Name

VK\_EXT\_validation\_cache - device extension



## [](#_registered_extension_number)Registered Extension Number

161

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Cort Stratton [\[GitHub\]cdwfs](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_validation_cache%5D%20%40cdwfs%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_validation_cache%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-08-29

**IP Status**

No known IP claims.

**Contributors**

- Cort Stratton, Google
- Chris Forbes, Google

## [](#_description)Description

This extension provides a mechanism for caching the results of potentially expensive internal validation operations across multiple runs of a Vulkan application. At the core is the [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html) object type, which is managed similarly to the existing [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html).

The new structure [VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html) can be included in the `pNext` chain at [vkCreateShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateShaderModule.html) time. It contains a [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html) to use when validating the [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html).

## [](#_new_object_types)New Object Types

- [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html)

## [](#_new_commands)New Commands

- [vkCreateValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateValidationCacheEXT.html)
- [vkDestroyValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyValidationCacheEXT.html)
- [vkGetValidationCacheDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetValidationCacheDataEXT.html)
- [vkMergeValidationCachesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMergeValidationCachesEXT.html)

## [](#_new_structures)New Structures

- [VkValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheCreateInfoEXT.html)
- Extending [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html), [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html):
  
  - [VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html)

## [](#_new_enums)New Enums

- [VkValidationCacheHeaderVersionEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheHeaderVersionEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkValidationCacheCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheCreateFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_VALIDATION_CACHE_EXTENSION_NAME`
- `VK_EXT_VALIDATION_CACHE_SPEC_VERSION`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_VALIDATION_CACHE_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_SHADER_MODULE_VALIDATION_CACHE_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_VALIDATION_CACHE_CREATE_INFO_EXT`

## [](#_version_history)Version History

- Revision 1, 2017-08-29 (Cort Stratton)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_validation_cache)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0