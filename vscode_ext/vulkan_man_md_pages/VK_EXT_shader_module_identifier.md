# VK_EXT_shader_module_identifier(3) Manual Page

## Name

VK_EXT_shader_module_identifier - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

463

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

        
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Version 1.1](#versions-1.1)  
     and  
    
[VK_EXT_pipeline_creation_cache_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_creation_cache_control.html)  
or  
[Version 1.3](#versions-1.3)  

## <a href="#_contact" class="anchor"></a>Contact

- Hans-Kristian Arntzen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_module_identifier%5D%20@HansKristian-Work%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_module_identifier%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>HansKristian-Work</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_shader_module_identifier](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_shader_module_identifier.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

Some applications generate SPIR-V code at runtime. When pipeline caches
are primed, either explicitly through e.g.
[VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html) mechanisms, or implicitly
through driver managed caches, having to re-generate SPIR-V modules is
redundant. SPIR-V modules could be cached on disk by an application, but
the extra disk size requirement might be prohibitive in some use cases.

This extension adds the ability for an application to query a small
identifier associated with a [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html). On
subsequent runs of the application, the same identifier **can** be
provided in lieu of a [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) object. A
pipeline creation call with such a module **may** succeed if a pipeline
could be created without invoking compilation, and information inside
the SPIR-V module is not required by the implementation.

`VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` **must** be
used if only the identifier is provided, and this use case is intended
to work like a non-blocking, speculative compile. Applications **can**
fallback as necessary.

The main motivation for identifying the module itself and not the entire
pipeline is that pipeline identifiers change when a driver is updated,
but module identifiers are expected to be stable for any particular
driver implementation. This approach is helpful for shader
pre-compilation systems which can prime pipeline caches ahead of time.
When on-disk pipeline caches are updated, the same shader identifiers
could lead to a pipeline cache hit.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetShaderModuleCreateInfoIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderModuleCreateInfoIdentifierEXT.html)

- [vkGetShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderModuleIdentifierEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleIdentifierEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT.html)

- Extending
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html):

  - [VkPipelineShaderStageModuleIdentifierCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageModuleIdentifierCreateInfoEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SHADER_MODULE_IDENTIFIER_EXTENSION_NAME`

- `VK_EXT_SHADER_MODULE_IDENTIFIER_SPEC_VERSION`

- `VK_MAX_SHADER_MODULE_IDENTIFIER_SIZE_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_MODULE_IDENTIFIER_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_MODULE_IDENTIFIER_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_MODULE_IDENTIFIER_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_SHADER_MODULE_IDENTIFIER_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-03-16 (Hans-Kristian Arntzen)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_shader_module_identifier"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
