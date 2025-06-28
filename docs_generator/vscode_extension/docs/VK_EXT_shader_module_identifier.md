# VK\_EXT\_shader\_module\_identifier(3) Manual Page

## Name

VK\_EXT\_shader\_module\_identifier - device extension



## [](#_registered_extension_number)Registered Extension Number

463

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Vulkan Version 1.1](#versions-1.1)  
     and  
     [VK\_EXT\_pipeline\_creation\_cache\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_creation_cache_control.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_contact)Contact

- Hans-Kristian Arntzen [\[GitHub\]HansKristian-Work](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_module_identifier%5D%20%40HansKristian-Work%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_module_identifier%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_shader\_module\_identifier](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_shader_module_identifier.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-05-16

**IP Status**

No known IP claims.

**Contributors**

- Hans-Kristian Arntzen, Valve
- Ricardo Garcia, Igalia
- Piers Daniell, NVIDIA
- Jan-Harald Fredriksen, Arm
- Tom Olson, Arm
- Faith Ekstrand, Collabora

## [](#_description)Description

Some applications generate SPIR-V code at runtime. When pipeline caches are primed, either explicitly through e.g. [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) mechanisms, or implicitly through driver managed caches, having to re-generate SPIR-V modules is redundant. SPIR-V modules could be cached on disk by an application, but the extra disk size requirement might be prohibitive in some use cases.

This extension adds the ability for an application to query a small identifier associated with a [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html). On subsequent runs of the application, the same identifier **can** be provided in lieu of a [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html) object. A pipeline creation call with such a module **may** succeed if a pipeline could be created without invoking compilation, and information inside the SPIR-V module is not required by the implementation.

`VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` **must** be used if only the identifier is provided, and this use case is intended to work like a non-blocking, speculative compile. Applications **can** fallback as necessary.

The main motivation for identifying the module itself and not the entire pipeline is that pipeline identifiers change when a driver is updated, but module identifiers are expected to be stable for any particular driver implementation. This approach is helpful for shader pre-compilation systems which can prime pipeline caches ahead of time. When on-disk pipeline caches are updated, the same shader identifiers could lead to a pipeline cache hit.

## [](#_new_commands)New Commands

- [vkGetShaderModuleCreateInfoIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderModuleCreateInfoIdentifierEXT.html)
- [vkGetShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderModuleIdentifierEXT.html)

## [](#_new_structures)New Structures

- [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleIdentifierEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT.html)
- Extending [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html):
  
  - [VkPipelineShaderStageModuleIdentifierCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageModuleIdentifierCreateInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SHADER_MODULE_IDENTIFIER_EXTENSION_NAME`
- `VK_EXT_SHADER_MODULE_IDENTIFIER_SPEC_VERSION`
- `VK_MAX_SHADER_MODULE_IDENTIFIER_SIZE_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_MODULE_IDENTIFIER_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_MODULE_IDENTIFIER_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_MODULE_IDENTIFIER_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_SHADER_MODULE_IDENTIFIER_EXT`

## [](#_version_history)Version History

- Revision 1, 2022-03-16 (Hans-Kristian Arntzen)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_shader_module_identifier)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0