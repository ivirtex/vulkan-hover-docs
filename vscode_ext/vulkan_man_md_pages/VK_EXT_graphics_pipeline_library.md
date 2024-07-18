# VK_EXT_graphics_pipeline_library(3) Manual Page

## Name

VK_EXT_graphics_pipeline_library - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

321

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Version 1.1](#versions-1.1)  
and  
[VK_KHR_pipeline_library](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_library.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_graphics_pipeline_library%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_graphics_pipeline_library%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_graphics_pipeline_library](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_graphics_pipeline_library.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension allows the separate compilation of four distinct parts of
graphics pipelines, with the intent of allowing faster pipeline loading
for applications reusing the same shaders or state in multiple
pipelines. Each part can be independently compiled into a graphics
pipeline library, with a final link step required to create an
executable pipeline that can be bound to a command buffer.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html):

  - [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceGraphicsPipelineLibraryFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGraphicsPipelineLibraryFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkGraphicsPipelineLibraryFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryFlagBitsEXT.html)

- [VkPipelineLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateFlagBits.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkGraphicsPipelineLibraryFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_GRAPHICS_PIPELINE_LIBRARY_EXTENSION_NAME`

- `VK_EXT_GRAPHICS_PIPELINE_LIBRARY_SPEC_VERSION`

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT`

  - `VK_PIPELINE_CREATE_RETAIN_LINK_TIME_OPTIMIZATION_INFO_BIT_EXT`

- Extending
  [VkPipelineLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateFlagBits.html):

  - `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_LIBRARY_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GRAPHICS_PIPELINE_LIBRARY_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GRAPHICS_PIPELINE_LIBRARY_PROPERTIES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-08-17 (Tobias Hector)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_graphics_pipeline_library"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
