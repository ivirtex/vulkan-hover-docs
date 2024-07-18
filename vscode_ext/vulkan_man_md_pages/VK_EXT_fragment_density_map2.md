# VK_EXT_fragment_density_map2(3) Manual Page

## Name

VK_EXT_fragment_density_map2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

333

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_fragment_density_map](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_fragment_density_map2%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_fragment_density_map2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-06-16

**Interactions and External Dependencies**  
- Interacts with Vulkan 1.1

**Contributors**  
- Matthew Netsch, Qualcomm Technologies, Inc.

- Jonathan Tinkham, Qualcomm Technologies, Inc.

- Jonathan Wicks, Qualcomm Technologies, Inc.

- Jan-Harald Fredriksen, ARM

## <a href="#_description" class="anchor"></a>Description

This extension adds additional features and properties to
[`VK_EXT_fragment_density_map`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map.html) in
order to reduce fragment density map host latency as well as improved
queries for subsampled sampler implementation-dependent behavior.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceFragmentDensityMap2FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMap2FeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceFragmentDensityMap2PropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMap2PropertiesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_FRAGMENT_DENSITY_MAP_2_EXTENSION_NAME`

- `VK_EXT_FRAGMENT_DENSITY_MAP_2_SPEC_VERSION`

- Extending [VkImageViewCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateFlagBits.html):

  - `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_2_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_2_PROPERTIES_EXT`

## <a href="#_examples" class="anchor"></a>Examples

### <a href="#_additional_limits_for_subsampling" class="anchor"></a>Additional Limits for Subsampling

Some implementations may not support subsampled samplers if certain
implementation limits are not observed by the app.
[VkPhysicalDeviceFragmentDensityMap2PropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMap2PropertiesEXT.html)
provides additional limits to require apps remain within these
boundaries if they wish to use subsampling.

### <a href="#_improved_host_latency" class="anchor"></a>Improved Host Latency

By default, the fragment density map is locked by the host for reading
between [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html) during
recording and `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
during draw execution.

This can introduce large latency for certain use cases between recording
the frame and displaying the frame. Apps may wish to modify the fragment
density map just before draw execution.

`VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT` is intended
to help address this for implementations that do not support the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentDensityMapDynamic"
target="_blank"
rel="noopener"><code>fragmentDensityMapDynamic</code></a> feature by
deferring the start of the locked range to
[vkEndCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEndCommandBuffer.html).

``` c
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

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-06-16 (Matthew Netsch)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_fragment_density_map2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
