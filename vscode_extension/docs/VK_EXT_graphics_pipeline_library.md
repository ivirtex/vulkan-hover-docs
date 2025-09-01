# VK\_EXT\_graphics\_pipeline\_library(3) Manual Page

## Name

VK\_EXT\_graphics\_pipeline\_library - device extension



## [](#_registered_extension_number)Registered Extension Number

321

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_KHR\_pipeline\_library](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_library.html)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_graphics_pipeline_library%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_graphics_pipeline_library%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_graphics\_pipeline\_library](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_graphics_pipeline_library.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-08-17

**Contributors**

- Tobias Hector, AMD
- Chris Glover, Google
- Jeff Leger, Qualcomm
- Jan-Harald Fredriksen, Arm
- Piers Daniell, NVidia
- Boris Zanin, Mobica
- Krzysztof Niski, NVidia
- Dan Ginsburg, Valve
- Sebastian Aaltonen, Unity
- Arseny Kapoulkine, Roblox
- Calle Lejdfors, Ubisoft
- Tiago Rodrigues, Ubisoft
- Francois Duranleau, Gameloft

## [](#_description)Description

This extension allows the separate compilation of four distinct parts of graphics pipelines, with the intent of allowing faster pipeline loading for applications reusing the same shaders or state in multiple pipelines. Each part can be independently compiled into a graphics pipeline library, with a final link step required to create an executable pipeline that can be bound to a command buffer.

## [](#_new_structures)New Structures

- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html):
  
  - [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceGraphicsPipelineLibraryFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceGraphicsPipelineLibraryFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT.html)

## [](#_new_enums)New Enums

- [VkGraphicsPipelineLibraryFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineLibraryFlagBitsEXT.html)
- [VkPipelineLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateFlagBits.html)

## [](#_new_bitmasks)New Bitmasks

- [VkGraphicsPipelineLibraryFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineLibraryFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_GRAPHICS_PIPELINE_LIBRARY_EXTENSION_NAME`
- `VK_EXT_GRAPHICS_PIPELINE_LIBRARY_SPEC_VERSION`
- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html):
  
  - `VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT`
  - `VK_PIPELINE_CREATE_RETAIN_LINK_TIME_OPTIMIZATION_INFO_BIT_EXT`
- Extending [VkPipelineLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateFlagBits.html):
  
  - `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_LIBRARY_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GRAPHICS_PIPELINE_LIBRARY_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GRAPHICS_PIPELINE_LIBRARY_PROPERTIES_EXT`

## [](#_version_history)Version History

- Revision 1, 2021-08-17 (Tobias Hector)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_graphics_pipeline_library).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0