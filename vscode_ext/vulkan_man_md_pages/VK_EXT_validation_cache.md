# VK_EXT_validation_cache(3) Manual Page

## Name

VK_EXT_validation_cache - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

161

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Cort Stratton <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_validation_cache%5D%20@cdwfs%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_validation_cache%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cdwfs</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-08-29

**IP Status**  
No known IP claims.

**Contributors**  
- Cort Stratton, Google

- Chris Forbes, Google

## <a href="#_description" class="anchor"></a>Description

This extension provides a mechanism for caching the results of
potentially expensive internal validation operations across multiple
runs of a Vulkan application. At the core is the
[VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html) object type, which is
managed similarly to the existing
[VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html).

The new struct
[VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html)
can be included in the `pNext` chain at
[vkCreateShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateShaderModule.html) time. It contains a
[VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html) to use when validating
the [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html).

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateValidationCacheEXT.html)

- [vkDestroyValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyValidationCacheEXT.html)

- [vkGetValidationCacheDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetValidationCacheDataEXT.html)

- [vkMergeValidationCachesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMergeValidationCachesEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheCreateInfoEXT.html)

- Extending [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html),
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html):

  - [VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkValidationCacheHeaderVersionEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheHeaderVersionEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkValidationCacheCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheCreateFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_VALIDATION_CACHE_EXTENSION_NAME`

- `VK_EXT_VALIDATION_CACHE_SPEC_VERSION`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_VALIDATION_CACHE_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_SHADER_MODULE_VALIDATION_CACHE_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_VALIDATION_CACHE_CREATE_INFO_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-08-29 (Cort Stratton)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_validation_cache"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
