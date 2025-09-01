# VK\_SEC\_pipeline\_cache\_incremental\_mode(3) Manual Page

## Name

VK\_SEC\_pipeline\_cache\_incremental\_mode - device extension



## [](#_registered_extension_number)Registered Extension Number

638

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Chris Hambacher [\[GitHub\]chambacher](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_SEC_pipeline_cache_incremental_mode%5D%20%40chambacher%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_SEC_pipeline_cache_incremental_mode%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-06-24

**IP Status**

No known IP claims.

**Contributors**

- Chris Hambacher, Samsung
- Mohan Maiya, Samsung
- Brandon Schade, Samsung

## [](#_description)Description

This extension allows layered implementations such as ANGLE to modify the default behavior of VkPipelineCache to return only the incremental data from the previous call to vkGetPipelineCacheData. Application developers should avoid using this extension.

Note

There is currently no specification language written for this extension. The links to APIs defined by the extension are to stubs that only include generated content such as API declarations and implicit valid usage statements.

Note

This extension is only intended for use in specific embedded environments with known implementation details, and is therefore undocumented.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePipelineCacheIncrementalModeFeaturesSEC](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineCacheIncrementalModeFeaturesSEC.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_SEC_PIPELINE_CACHE_INCREMENTAL_MODE_EXTENSION_NAME`
- `VK_SEC_PIPELINE_CACHE_INCREMENTAL_MODE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_CACHE_INCREMENTAL_MODE_FEATURES_SEC`

## [](#_version_history)Version History

- Revision 1, 2025-06-24 (Chris Hambacher)
  
  - Initial specification

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_SEC_pipeline_cache_incremental_mode).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0