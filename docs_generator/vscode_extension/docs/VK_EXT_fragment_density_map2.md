# VK\_EXT\_fragment\_density\_map2(3) Manual Page

## Name

VK\_EXT\_fragment\_density\_map2 - device extension



## [](#_registered_extension_number)Registered Extension Number

333

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_fragment\_density\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map.html)

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_fragment_density_map2%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_fragment_density_map2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-06-16

**Interactions and External Dependencies**

- Interacts with Vulkan 1.1

**Contributors**

- Matthew Netsch, Qualcomm Technologies, Inc.
- Jonathan Tinkham, Qualcomm Technologies, Inc.
- Jonathan Wicks, Qualcomm Technologies, Inc.
- Jan-Harald Fredriksen, ARM

## [](#_description)Description

This extension adds additional features and properties to `VK_EXT_fragment_density_map` in order to reduce fragment density map host latency as well as improved queries for subsampled sampler implementation-dependent behavior.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceFragmentDensityMap2FeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentDensityMap2FeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceFragmentDensityMap2PropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentDensityMap2PropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_FRAGMENT_DENSITY_MAP_2_EXTENSION_NAME`
- `VK_EXT_FRAGMENT_DENSITY_MAP_2_SPEC_VERSION`
- Extending [VkImageViewCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateFlagBits.html):
  
  - `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_2_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_2_PROPERTIES_EXT`

## [](#_examples)Examples

### [](#_additional_limits_for_subsampling)Additional Limits for Subsampling

Some implementations may not support subsampled samplers if certain implementation limits are not observed by the app. [VkPhysicalDeviceFragmentDensityMap2PropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentDensityMap2PropertiesEXT.html) provides additional limits to require apps remain within these boundaries if they wish to use subsampling.

### [](#_improved_host_latency)Improved Host Latency

By default, the fragment density map is locked by the host for reading between [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderPass.html) during recording and `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT` during draw execution.

This can introduce large latency for certain use cases between recording the frame and displaying the frame. Apps may wish to modify the fragment density map just before draw execution.

`VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT` is intended to help address this for implementations that do not support the [`fragmentDensityMapDynamic`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-fragmentDensityMapDynamic) feature by deferring the start of the locked range to [vkEndCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEndCommandBuffer.html).

```c++
// Create view for fragment density map
VkImageViewCreateInfo createInfo =
{
   .sType = VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO,
   // ...
   .viewType = VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT,
   .format = fdmImage, // Created with VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT
   // ...
};

// ...

// Begin fragment density map render pass with deferred view.
// By default, fdmImage must not be modified after this call
// unless view is created with the FDM dynamic or deferred flags.
vkCmdBeginRenderPass(commandBuffer, pRenderPassBegin, contents);

// ...

vkCmdEndRenderPass(VkCommandBuffer commandBuffer);

// Can keep making modifications to deferred fragment density maps
// while recording commandBuffer ...

result = vkEndCommandBuffer(commandBuffer);

// Must now freeze modifying fdmImage until after the
// VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT of
// the last executing draw that uses the fdmImage in the
// last submit of commandBuffer.
```

## [](#_version_history)Version History

- Revision 1, 2020-06-16 (Matthew Netsch)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_fragment_density_map2)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0